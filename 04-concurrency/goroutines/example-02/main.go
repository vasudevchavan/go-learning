package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

// example1
func updateMessage(s string) {
	defer wg.Done()
	msg = s
}

// end

// func updateMessage(s string, m *sync.Mutex) {
// 	defer wg.Done()
// 	m.Lock()
// 	msg = s
// 	m.Unlock()
// }

func main() {
	msg = "hello"
	// Example 1
	// Start
	// execute using this command go run -race .
	// Problem below code will fail with DATA RACE warning.
	// Concurrent go routines are trying to access the same data
	wg.Add(2)
	go updateMessage("bangalore")
	go updateMessage("mysore")
	wg.Wait()
	fmt.Println(msg)
	// End

	// Example 2 to solve DATA RACE using mutex
	// Start
	// var mtex sync.Mutex
	// wg.Add(2)
	// go updateMessage("bangalore", &mtex)
	// go updateMessage("mysore", &mtex)
	// wg.Wait()
	// fmt.Println(msg)
	// End

}
