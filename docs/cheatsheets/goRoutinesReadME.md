✅ STAGE 1: Goroutines – Running Code Concurrently
🔹 What is a Goroutine?

A goroutine is a lightweight thread managed by Go. You create one with the go keyword.

🔸 Real-time Example: Background Logger

🧠 Imagine you're writing to a log file while your program is doing other things.

package main

import (
    "fmt"
    "time"
)

func logActivity() {
    for i := 1; i <= 5; i++ {
        fmt.Println("Logging activity", i)
        time.Sleep(500 * time.Millisecond)
    }
}

func main() {
    go logActivity() // run logActivity in the background

    fmt.Println("Doing main work")
    time.Sleep(3 * time.Second)
    fmt.Println("Main work done")
}

✅ STAGE 2: Basic Channels – Communication Between Goroutines
🔹 What is a Channel?

A channel is a way to send and receive data safely between goroutines.

🔸 Real-time Example: Sending a Message to a Worker
package main

import (
    "fmt"
)

func worker(ch chan string) {
    msg := <-ch
    fmt.Println("Worker received:", msg)
}

func main() {
    ch := make(chan string)

    go worker(ch)

    ch <- "Start processing job"
}


🧠 ch <- "..." sends data into the channel, <-ch receives it.

✅ STAGE 3: Buffered Channels & Closing
🔹 Buffered Channel

You can make a channel with a buffer to hold data temporarily.

func main() {
    ch := make(chan int, 2) // Can hold 2 values

    ch <- 1
    ch <- 2
    fmt.Println(<-ch)
    fmt.Println(<-ch)
}

🔸 Real-time Example: Queuing Emails
func sendEmail(email string, ch chan string) {
    ch <- "Email sent to: " + email
}

func main() {
    ch := make(chan string, 3)

    go sendEmail("alice@example.com", ch)
    go sendEmail("bob@example.com", ch)
    go sendEmail("carol@example.com", ch)

    fmt.Println(<-ch)
    fmt.Println(<-ch)
    fmt.Println(<-ch)
}

🔹 Closing a Channel

Close a channel when no more values will be sent. This allows receivers to know when to stop.

func main() {
    ch := make(chan int)

    go func() {
        for i := 1; i <= 3; i++ {
            ch <- i
        }
        close(ch)
    }()

    for val := range ch {
        fmt.Println("Received:", val)
    }
}

✅ STAGE 4: WaitGroup – Wait for Goroutines to Finish
🔸 Real-time Example: File Downloads
package main

import (
    "fmt"
    "sync"
    "time"
)

func downloadFile(filename string, wg *sync.WaitGroup) {
    defer wg.Done()
    fmt.Println("Downloading:", filename)
    time.Sleep(1 * time.Second)
    fmt.Println("Finished:", filename)
}

func main() {
    var wg sync.WaitGroup

    files := []string{"file1.txt", "file2.txt", "file3.txt"}

    for _, f := range files {
        wg.Add(1)
        go downloadFile(f, &wg)
    }

    wg.Wait()
    fmt.Println("All downloads complete")
}


🧠 WaitGroup helps you wait for all background work to finish before exiting.

✅ STAGE 5: Select Statement – Listening to Multiple Channels
🔸 Real-time Example: Handling API and Timeout
func main() {
    apiChan := make(chan string)
    timeout := time.After(2 * time.Second)

    go func() {
        time.Sleep(1 * time.Second)
        apiChan <- "API response"
    }()

    select {
    case msg := <-apiChan:
        fmt.Println("Received:", msg)
    case <-timeout:
        fmt.Println("Request timed out")
    }
}


🧠 select waits on multiple channel operations and picks the one that's ready first.

✅ STAGE 6: Combine It All
🔸 Real-time Mini Project: Task Queue with Worker Pool
package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
    defer wg.Done()
    for job := range jobs {
        fmt.Printf("Worker %d processing job %d\n", id, job)
        time.Sleep(500 * time.Millisecond)
    }
}

func main() {
    const numWorkers = 3
    const numJobs = 5

    jobs := make(chan int)
    var wg sync.WaitGroup

    // Start workers
    for w := 1; w <= numWorkers; w++ {
        wg.Add(1)
        go worker(w, jobs, &wg)
    }

    // Send jobs
    for j := 1; j <= numJobs; j++ {
        jobs <- j
    }
    close(jobs)

    wg.Wait()
    fmt.Println("All jobs done")
}

✅ RECAP – Concepts and Use Cases
Concept	Real Use Case
go func()	Run a task in background (logs, API)
Channels	Send results to main thread
Buffered Ch	Task queues
WaitGroup	Wait for all goroutines to finish
select	Listen to multiple inputs (timeout, signal)
🎯 Your Learning Plan Summary
Week 1

✅ Learn goroutines and basic channels

🧪 Build logger and ping-pong example

Week 2

✅ Learn buffered channels, closing, WaitGroup

🧪 Build file downloader with WaitGroup

Week 3

✅ Learn select, timeout, cancellation

🧪 Build task queue with worker pool