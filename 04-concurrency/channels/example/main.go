package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {

	// Example one
	// Buffered channels will never give error but Unbuffered Channel does.
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	// Example two
	// Buffered channels will never give error but Unbuffered Channel does.
	ch1 := make(chan []int, 3)
	ch1 <- []int{1}
	ch1 <- []int{2}
	ch1 <- []int{3}
	fmt.Print(<-ch1)
	fmt.Print(<-ch1)
	fmt.Print(<-ch1)

	// Example three
	// Buffered channels will never give error but Unbuffered Channel does.
	ch2 := make(chan []int, 10)
	ch2 <- []int{1}
	ch2 <- []int{2}
	ch2 <- []int{3}

	close(ch2)
	for i := range ch2 {
		fmt.Println(i)
	}

	// Example four
	// Buffered channels will never give error but Unbuffered Channel does.
	pchan := make(chan Person, 1)
	p := Person{Name: "Ramesh", Age: 10}
	pchan <- p
	fmt.Println(<-pchan)

	// Create functions to send and receive values from channels
	// go Routines for Sendint and Receiveint 
	n := make(chan int)
	go Sendint(n)
	go Receiveint(n)
	fmt.Println("\nWill not receive the values send to channel") 
	fmt.Println("Since Sendint & Receiveint will not return and Main function will not wait\n")

  // Values is received back sonce Receiveint function will block and wait for Sendint to send value.
  go Sendint(n)
	Receiveint(n)
	
}

func Sendint(n chan <- int) {
	n <- 1
}
func Receiveint( n <-chan int) {
	fmt.Println("value received from channel:",<-n)
	
}
