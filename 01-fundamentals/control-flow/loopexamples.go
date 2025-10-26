package main

import (
	"fmt"
	"math/rand"
)

func main() {

	// Sample if else statement
	testVal := rand.Intn(200)
	DemoIfElse(testVal)	// Demo if else statement
	DemoSwitch(testVal)	// Demo Switch statement
	DemoSelectUsingChannel()	// Demo select using channel
	DemoForStatement(testVal)	// Demo For statement
	OddNumber(testVal)	// Demo find odd number with for,if and Module Operator
}


func OddNumber(t int){
  cnt:=0
  for i:=1;i<t;i++{
    if i%2 == 1 {
      // fmt.Printf("Odd number:%d\n",i)
      cnt++
    }
  }
  fmt.Printf("found %d Odd numbers within Range %d",cnt,t)
}


func DemoForStatement(t int) {
	for i := 0; i < t; i++ {
		if i == 50 {
			fmt.Println("Values is 50")
		}
	}

  var t1 int
  t1 = t
	for t1 < 100 {
    fmt.Println("Inside second for loop")
    t1=t1+50
	}
}

func DemoIfElse(t int) {
	if t < 50 {
		fmt.Println("Values within 50")
	} else if t > 50 && t < 100 {
		fmt.Println("Value in between 50 && 100")
	} else if t > 100 && t < 200 {
		fmt.Println("Value in between 100 && 200")
	}
}

func DemoSwitch(t int) {

	switch {
	case t < 50:
		fmt.Println("Values within 50")
	case t > 50 && t < 100:
		fmt.Println("Value in between 50 && 100")
	case t > 100 && t < 200:
		fmt.Println("Value in between 100 && 200")
	default:
		fmt.Println("Try once again")
	}
}

func DemoSelectUsingChannel() {
	chn1 := make(chan int)
	chn2 := make(chan int)

	go func() {
		chn1 <- 45
	}()

	go func() {
		chn1 <- 47
	}()

	select {
	case test := <-chn1:
		fmt.Printf("Received %d on channel1 \n", test)
	case test := <-chn2:
		fmt.Printf("Received %d on channel2\n", test)
	}

}
