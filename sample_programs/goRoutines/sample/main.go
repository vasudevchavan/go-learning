package main

import (
	"fmt"
	"sync"
)

func printGreeting(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}

func main() {

	// Declaration of wait group
	var wg sync.WaitGroup

	go printGreeting("hello", &wg)
	// time.Sleep(1 * time.Second) added comment because we added the waitgroup
	go printGreeting("go", &wg)
	// time.Sleep(1 * time.Second)
	go printGreeting("routines", &wg)
	// time.Sleep(1 * time.Second)
	go printGreeting("world", &wg)
	// time.Sleep(1 * time.Second)

	table := []string{
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
	}

	// Number decideded based on length of the string and explore each case
	// To understand how WaitGroup works
	wg.Add(len(table) + 4)
	// wg.Add(len(table) -1 )
	// wg.Add(len(table))
	// wg.Add(len(table) + 9)

	for i, v := range table {
		go printGreeting(fmt.Sprintf("we are going to print index %d %s", i, v), &wg)
	}
	// time.Sleep(1 * time.Second)

	wg.Wait()
}
