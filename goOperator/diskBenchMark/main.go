package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"time"
)

const (
	blockSize  = 1024 * 1024 // 1MB
	numThreads = 1           // one per type
)

type BenchmarkData struct {
	sync.Mutex
	BytesWritten int64
	Speeds       []float64
	CurrentTime  int
	Name         string
}

func (b *BenchmarkData) recordSpeed() {
	b.Lock()
	defer b.Unlock()
	mbps := float64(b.BytesWritten) / (1024 * 1024)
	b.Speeds = append(b.Speeds, mbps)
	if len(b.Speeds) > 120 {
		b.Speeds = b.Speeds[1:]
	}
	b.BytesWritten = 0
	b.CurrentTime++
}

func (b *BenchmarkData) stats() (avg, peak, stddev float64, points []DataPoint) {
	b.Lock()
	defer b.Unlock()
	startTime := b.CurrentTime - len(b.Speeds)
	var sum, sumSq float64
	for i, s := range b.Speeds {
		sum += s
		sumSq += s * s
		points = append(points, DataPoint{
			Time:  startTime + i,
			Speed: s,
		})
		if s > peak {
			peak = s
		}
	}
	n := float64(len(b.Speeds))
	if n > 0 {
		avg = sum / n
		stddev = math.Sqrt((sumSq / n) - (avg * avg))
	}
	return
}

type DataPoint struct {
	Time  int     `json:"time"`
	Speed float64 `json:"speed"`
}

var (
	seqWriteData  = &BenchmarkData{Name: "Sequential Write"}
	seqReadData   = &BenchmarkData{Name: "Sequential Read"}
	randWriteData = &BenchmarkData{Name: "Random Write"}
	randReadData  = &BenchmarkData{Name: "Random Read"}
)

func workerSequentialWrite(b *BenchmarkData, done <-chan struct{}) {
	filename := "seq_write.tmp"
	file, err := os.Create(filename)
	if err != nil {
		log.Println("Error:", err)
		return
	}
	defer os.Remove(filename)
	defer file.Close()

	data := make([]byte, blockSize)

	for {
		select {
		case <-done:
			return
		default:
			n, err := file.Write(data)
			if err != nil {
				log.Println("Write error:", err)
				return
			}
			b.Lock()
			b.BytesWritten += int64(n)
			b.Unlock()
		}
	}
}

func workerSequentialRead(b *BenchmarkData, done <-chan struct{}) {
	filename := "seq_read.tmp"
	data := make([]byte, blockSize)
	err := os.WriteFile(filename, make([]byte, 1024*1024*500), 0644) // 500MB
	if err != nil {
		log.Println("Error creating file:", err)
		return
	}
	defer os.Remove(filename)

	file, err := os.Open(filename)
	if err != nil {
		log.Println("Open error:", err)
		return
	}
	defer file.Close()

	for {
		select {
		case <-done:
			return
		default:
			n, err := file.Read(data)
			if err != nil {
				file.Seek(0, 0) // Loop
				continue
			}
			b.Lock()
			b.BytesWritten += int64(n)
			b.Unlock()
		}
	}
}

func workerRandomWrite(b *BenchmarkData, done <-chan struct{}) {
	filename := "rand_write.tmp"
	size := 1024 * 1024 * 500 // 500MB
	file, err := os.Create(filename)
	if err != nil {
		log.Println("Error:", err)
		return
	}
	defer os.Remove(filename)
	defer file.Close()

	file.Write(make([]byte, size)) // preallocate
	data := make([]byte, blockSize)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for {
		select {
		case <-done:
			return
		default:
			offset := int64(r.Intn(size/blockSize) * blockSize)
			file.Seek(offset, 0)
			n, err := file.Write(data)
			if err != nil {
				log.Println("Write error:", err)
				return
			}
			b.Lock()
			b.BytesWritten += int64(n)
			b.Unlock()
		}
	}
}

func workerRandomRead(b *BenchmarkData, done <-chan struct{}) {
	filename := "rand_read.tmp"
	size := 1024 * 1024 * 500 // 500MB
	err := os.WriteFile(filename, make([]byte, size), 0644)
	if err != nil {
		log.Println("Error:", err)
		return
	}
	defer os.Remove(filename)

	file, err := os.Open(filename)
	if err != nil {
		log.Println("Open error:", err)
		return
	}
	defer file.Close()

	data := make([]byte, blockSize)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for {
		select {
		case <-done:
			return
		default:
			offset := int64(r.Intn(size/blockSize) * blockSize)
			file.Seek(offset, 0)
			n, err := file.Read(data)
			if err != nil {
				continue
			}
			b.Lock()
			b.BytesWritten += int64(n)
			b.Unlock()
		}
	}
}

func startSampler(done <-chan struct{}) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			seqWriteData.recordSpeed()
			seqReadData.recordSpeed()
			randWriteData.recordSpeed()
			randReadData.recordSpeed()
		}
	}
}

func serveWeb() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, htmlPage)
	})

	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {

		type ChartData struct {
			Label   string      `json:"label"`
			Average float64     `json:"average"`
			Peak    float64     `json:"peak"`
			StdDev  float64     `json:"stddev"`
			Points  []DataPoint `json:"points"`
		}
		charts := map[string]ChartData{
			"seqWrite":  toChartData(seqWriteData),
			"seqRead":   toChartData(seqReadData),
			"randWrite": toChartData(randWriteData),
			"randRead":  toChartData(randReadData),
		}
		json.NewEncoder(w).Encode(charts)
	})

	http.HandleFunc("/export", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Disposition", "attachment; filename=\"benchmark.csv\"")
		w.Header().Set("Content-Type", "text/csv")
		writer := csv.NewWriter(w)
		writer.Write([]string{"Type", "Time(s)", "Speed(MB/s)"})
		writeCSV(writer, "Sequential Write", seqWriteData)
		writeCSV(writer, "Sequential Read", seqReadData)
		writeCSV(writer, "Random Write", randWriteData)
		writeCSV(writer, "Random Read", randReadData)
		writer.Flush()
	})

	fmt.Println("🌐 Open http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func toChartData(b *BenchmarkData) (data struct {
	Label   string      `json:"label"`
	Average float64     `json:"average"`
	Peak    float64     `json:"peak"`
	StdDev  float64     `json:"stddev"`
	Points  []DataPoint `json:"points"`
}) {
	data.Label = b.Name
	data.Average, data.Peak, data.StdDev, data.Points = b.stats()
	return
}

func writeCSV(w *csv.Writer, label string, b *BenchmarkData) {
	b.Lock()
	defer b.Unlock()
	start := b.CurrentTime - len(b.Speeds)
	for i, s := range b.Speeds {
		w.Write([]string{
			label,
			strconv.Itoa(start + i),
			fmt.Sprintf("%.2f", s),
		})
	}
}

const htmlPage = `
<!DOCTYPE html>
<html>
<head>
	<title>Disk Benchmark Dashboard</title>
	<script src="https://cdn.jsdelivr.net/npm/chart.js"></script>
	<style>
		body { font-family: sans-serif; margin: 20px; background: #f5f5f5; }
		h2 { text-align: center; }
		.grid {
    display: grid;
    grid-template-columns: repeat(4, 1fr); /* 4 columns */
    gap: 20px;
}

.chart-container {
    background: white;
    padding: 10px;
    border-radius: 8px;
    box-shadow: 0 0 10px rgba(0,0,0,0.1);
    width: 100%;   /* full width of grid cell */
    height: 300px; /* fixed height for charts */
    display: flex;
    flex-direction: column;
}

.chart-container canvas {
    flex-grow: 1; /* make chart fill available vertical space */
}
		.stats {
			margin-top: 10px;
			font-size: 14px;
		}
		#exportBtn {
			margin: 20px auto;
			display: block;
			padding: 10px 20px;
			background: #007bff;
			color: white;
			border: none;
			border-radius: 5px;
			cursor: pointer;
		}
	</style>
</head>
<body>
	<h2>Disk I/O Benchmark Dashboard</h2>
	<button id="exportBtn" onclick="downloadCSV()">⬇️ Download CSV</button>
	<div class="grid">
		<div class="chart-container">
			<h3>Sequential Write</h3>
			<canvas id="seqWriteChart"></canvas>
			<div class="stats" id="seqWriteStats"></div>
		</div>
		<div class="chart-container">
			<h3>Sequential Read</h3>
			<canvas id="seqReadChart"></canvas>
			<div class="stats" id="seqReadStats"></div>
		</div>
		<div class="chart-container">
			<h3>Random Write</h3>
			<canvas id="randWriteChart"></canvas>
			<div class="stats" id="randWriteStats"></div>
		</div>
		<div class="chart-container">
			<h3>Random Read</h3>
			<canvas id="randReadChart"></canvas>
			<div class="stats" id="randReadStats"></div>
		</div>
	</div>

	<script>
		const charts = {
			seqWrite: setupChart('seqWriteChart'),
			seqRead: setupChart('seqReadChart'),
			randWrite: setupChart('randWriteChart'),
			randRead: setupChart('randReadChart')
		};

		function setupChart(id) {
			const ctx = document.getElementById(id).getContext('2d');
			return new Chart(ctx, {
				type: 'line',
				data: {
					labels: [],
					datasets: [{
						label: 'Speed (MB/s)',
						data: [],
						borderColor: [],
						backgroundColor: [],
						fill: false,
						tension: 0.2
					}]
				},
				options: {
					scales: {
						x: {
							title: { display: true, text: 'Time (s)' }
						},
						y: {
							title: { display: true, text: 'MB/s' },
							beginAtZero: true
						}
					},
					animation: false,
					responsive: true,
					maintainAspectRatio: false
				}
			});
		}

		function getColor(speed) {
			if (speed > 80) return 'green';
			if (speed > 40) return 'orange';
			return 'red';
		}

		async function fetchData() {
			const res = await fetch('/data');
			const data = await res.json();

			for (const key in charts) {
				const chart = charts[key];
				const item = data[key];

				const labels = item.points.map(p => p.time);
				const speeds = item.points.map(p => p.speed);
				const colors = speeds.map(getColor);

				chart.data.labels = labels;
				chart.data.datasets[0].data = speeds;
				chart.data.datasets[0].borderColor = colors;
				chart.update();

				document.getElementById(key + 'Stats').innerHTML =
    				"<strong>Average:</strong> " + item.average.toFixed(2) + " MB/s<br>" +
    				"<strong>Peak:</strong> " + item.peak.toFixed(2) + " MB/s<br>" +
    				"<strong>Std Dev:</strong> " + item.stddev.toFixed(2) + " MB/s";
			}
		}

		function downloadCSV() {
			window.location.href = "/export";
		}

		setInterval(fetchData, 1000);
	</script>
</body>
</html>
`

func main() {
	done := make(chan struct{})
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	go func() {
		<-signalChan
		fmt.Println("\n🛑 Stopping benchmarks...")
		close(done)
	}()

	go workerSequentialWrite(seqWriteData, done)
	go workerSequentialRead(seqReadData, done)
	go workerRandomWrite(randWriteData, done)
	go workerRandomRead(randReadData, done)
	go startSampler(done)

	serveWeb()
}
