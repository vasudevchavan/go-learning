package main

import (
	"fmt"
	"sync"
	"time"
)


func sendMsg(wg *sync.WaitGroup, msg string, in int) {
	if wg != nil {
	defer wg.Done()	
	}
	
	for i := 0; i < 2; i++ {
		time.Sleep(2 * time.Millisecond)
	}
	msg = "testing-end"
}

// Sample function without concurrent go Routines calls
func withConcurrencyGoRoutines() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ 	{
		wg.Add(1)
		
		go sendMsg(&wg, sendmessage, i)
	}
	wg.Wait()

}

// Sample function with concurrent go Routines calls
func withoutConcurrencyGoRoutines() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go sendMsg(&wg, sendmessage, i)
		wg.Wait()
	}
	for i := 0; i < 10; i++ {
		sendMsg(nil, sendmessage, i)
	}
}

func oddEvenNumber(innum []int) {
	var wg sync.WaitGroup

	fmt.Println("Even Number are")
	for _, num := range innum {
		wg.Add(1)
		num := num
		go func() {
			defer wg.Done()
			if num%2 == 0 {
				fmt.Println(num)
			}
		}()
	}
	wg.Wait()
}
