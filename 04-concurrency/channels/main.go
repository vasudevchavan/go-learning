package main

import (
	"fmt"
)

func main() {

	// Sample channel with single value sending and receiving it.
	singleChan := make(chan int)
	go func() {
		singleChan <- 45
	}()
	fmt.Println("We are receving a single value", <-singleChan)
	close(singleChan)

	// Sample example of channel
	fmt.Println("Sending and received bool")
	done := make(chan bool)
	go worker(done)
	<-done
	close(done)

	// Sample example to send and receive an Int
	fmt.Println("\nSend and Receive single int")
	singleInt := make(chan int)
	go receiveInt(singleInt)
	fmt.Printf("We received singe integer %v \n", <-singleInt)
	// close(singleInt) Leds to error since this channel is closed inside goRoutine

	// Sample example to Print range of int
	fmt.Println("\nPrinting int using channels")
	numbers := make(chan int)
	go printNumbers(numbers)
	for i := range numbers {
		fmt.Println(i)
	}

	// Sample example to Print only even numbers
	fmt.Printf("\nPrinting even numbers")
	evenNumbers := make(chan int)
	go getEvenNumbers(evenNumbers, 6)
	for num := range evenNumbers {
		fmt.Printf("%d\n", num)
	}

	// Sample example to Print odd and even numbers using different channels

	fmt.Println("Printing Odd Even number sequentially ")
	getEven := make(chan int)
	getOdd := make(chan int)

	go getEvenNumbers(getEven, 20)
	go getOddNumbers(getOdd, 20)

	doneEven := make(chan bool)
	doneOdd := make(chan bool)

	go func() {
		for num := range getEven {
			fmt.Printf("Even:%d\n", num)
		}
		doneEven <- true
	}()

	go func() {
		for num := range getOdd {
			fmt.Printf("Odd:%d\n", num)
		}
		doneOdd <- true
	}()

	<-doneEven
	<-doneOdd

	// Sample example to Print odd and even numbers using different channels using different Method

	fmt.Println("Printing Odd Even number parallel")
	getEvenN := make(chan bool)
	getOddN := make(chan bool)

	limit := 20
	go getEvenNumbersN(getEvenN, getOddN, limit)
	go getOddNumbersN(getEvenN, getOddN, limit)

	getEvenN <- true

	for range getEvenN {
	}
	for range getOddN {
	}

	// Example of sequential and parallel channel usage

	pInt := make(chan int)

	fmt.Println("Sequential")
	go func() {
		for i := 0; i < 5; i++ {
			pInt <- i
		}
		close(pInt)
	}()

	for val := range pInt {
		fmt.Println(val)
	}

	p1Int := make(chan int)
	for i := 0; i < 5; i++ {
		go func(num int) {
			p1Int <- num
		}(i)
	}

	fmt.Println("Prallel")
	for i := 0; i < 5; i++ {
		fmt.Println(<-p1Int)
	}
}
