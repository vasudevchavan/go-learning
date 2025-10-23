package main

import (
	"fmt"
	"time"
)

func sendMsg(msg string, in int, en bool) {

	if en {
		defer wg.Done()
		for i := 0; i < 2; i++ {
			time.Sleep(2 * time.Millisecond)
		}
	}
	msg = "testing-end"
}

// Sample function without concurrent go Routines calls
func withConcurrencyGoRoutines() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		sendmessage = "go routine testing"
		go sendMsg(sendmessage, i, true)
	}
	wg.Wait()
	for i := 0; i < 10; i++ {
		sendmessage = "Normal testing"
		sendMsg(sendmessage, i, false)
	}
}

// Sample function with concurrent go Routines calls
func withoutConcurrencyGoRoutines() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		sendmessage = "go routine testing"
		go sendMsg(sendmessage, i, true)
		wg.Wait()
	}
	for i := 0; i < 10; i++ {
		sendmessage = "Normal testing"
		sendMsg(sendmessage, i, false)
	}
}

func oddEvenNumber(innum []int) {
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
