package main

import (
	"fmt"
	"log"
	"math"
	// "os"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// Range of n values
	var nValues []float64
	for i := 1; i <= 10; i++ {
		nValues = append(nValues, float64(i))
	}

	// Compute different Big O growths
	O1 := make(plotter.XYs, len(nValues))
	OlogN := make(plotter.XYs, len(nValues))
	On := make(plotter.XYs, len(nValues))
	OnLogN := make(plotter.XYs, len(nValues))
	On2 := make(plotter.XYs, len(nValues))
	O2n := make(plotter.XYs, len(nValues))
	OnFact := make(plotter.XYs, len(nValues))

	for i, n := range nValues {
		O1[i].X, O1[i].Y = n, 1
		OlogN[i].X, OlogN[i].Y = n, math.Log2(n)
		On[i].X, On[i].Y = n, n
		OnLogN[i].X, OnLogN[i].Y = n, n*math.Log2(n)
		On2[i].X, On2[i].Y = n, n*n
		O2n[i].X, O2n[i].Y = n, math.Pow(2, n)
		OnFact[i].X, OnFact[i].Y = n, factorial(n)
	}

	// Create a new plot
	p := plot.New()
	p.Title.Text = "Growth of Big O Time Complexities"
	p.X.Label.Text = "n (input size)"
	p.Y.Label.Text = "Operations (relative scale)"

	// Add lines to plot
	addLine(p, O1, "O(1)")
	addLine(p, OlogN, "O(log n)")
	addLine(p, On, "O(n)")
	addLine(p, OnLogN, "O(n log n)")
	addLine(p, On2, "O(n²)")
	addLine(p, O2n, "O(2ⁿ)")
	addLine(p, OnFact, "O(n!)")

	p.Legend.Top = true

	// Limit Y to make the graph readable
	p.Y.Min = 0
	p.Y.Max = 1000

	// Save the plot as an image
	if err := p.Save(6*vg.Inch, 4*vg.Inch, "big_o_growth.png"); err != nil {
		log.Fatalf("Failed to save plot: %v", err)
	}

	fmt.Println("✅ Chart saved as big_o_growth.png")
}

func factorial(n float64) float64 {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

func addLine(p *plot.Plot, pts plotter.XYs, label string) {
	line, err := plotter.NewLine(pts)
	if err != nil {
		log.Fatalf("Failed to create line: %v", err)
	}
	line.LineStyle.Width = vg.Points(1)
	p.Add(line)
	p.Legend.Add(label, line)
}
