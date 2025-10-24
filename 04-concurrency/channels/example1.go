package main

import (
	"fmt"
	"time"
)

func worker(d chan bool) {
	fmt.Println("Processing")
	time.Sleep(2 * time.Second)
	fmt.Println("Completed")
	d <- false
}

func receiveInt(sendInt chan int) {
	sendInt <- 44
	close(sendInt)
}

func printNumbers(n chan int) {
	for i := range 10 {
		n <- i
	}
	close(n)
}

func getEvenNumbers(n chan int, limit int) {
	for i := range limit {
		if i%2 == 0 {
			n <- i
		}
	}
	close(n)
}

func getOddNumbers(n chan int, limit int) {
	for i := range limit {
		if i%2 != 0 {
			n <- i
		}
	}
	close(n)
}

func getEvenNumbersN(e chan bool, o chan bool, limit int) {
	<-o
	for i := range limit {
		if i%2 == 0 {
			fmt.Printf("Even:%d\n", i)
			e <- true
		}
	}
	close(e)

}

func getOddNumbersN(e chan bool, o chan bool, limit int) {
	<-e
	for i := range limit {
		if i%2 != 0 {
			fmt.Printf("Odd:%d\n", i)
			o <- true
		}
	}
	close(o)
}
