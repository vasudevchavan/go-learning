package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func printGreeting() {
	fmt.Println(msg)
}

func updateGreeting(s string) {
	defer wg.Done()
	msg = s
}

func main() {
	msg = "Hello go routines"

	wg.Add(1)
	go updateGreeting("Hello go")
	wg.Wait()
	printGreeting()
	wg.Add(1)
	go updateGreeting("Hello golang")
	wg.Wait()
	printGreeting()
	wg.Add(1)
	go updateGreeting("Hello goroutine")
	wg.Wait()
	printGreeting()

}
