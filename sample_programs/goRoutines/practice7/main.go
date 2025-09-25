package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan string) {
	for job := range jobs {
		fmt.Printf("Worker %d processing the %s Job \n", id, job)
	}
}

func main() {
	jobs := make(chan string)

	for i := 1; i <= 4; i++ {
		go worker(i, jobs)
	}

	for _, job := range []string{"a", "b", "c", "d", "e", "f", "g", "h"} {
		jobs <- job
		fmt.Println(job)
	}
	time.Sleep(1 * time.Second)
	close(jobs)
}
