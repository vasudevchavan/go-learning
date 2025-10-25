package main

import (
	"fmt"
	"os"
	"strings"
)

var debugFlag bool = false

func getPrimeCount(p int) int {

	// 0,1 are not Not Prime
	if p < 2 {
		return 0
	}

	// Creating a slnce of bool and declare every number is prime within range of p
	isPrime := make([]bool, p)
	for i := 2; i < p; i++ {
		isPrime[i] = true
	}

	//
	for i := 2; i*i < p; i++ {
		if debugFlag {
			fmt.Printf("Values of i %v before calculating # prime or not \n", i)
		}
		if isPrime[i] {
			for mul := i * i; mul < p; mul += i {
				isPrime[mul] = false
			}
		}
	}

	if debugFlag {
		for k, v := range isPrime {
			fmt.Printf("index %v and isPrime bool flag %v \n", k, v)
		}
		fmt.Println("Print list of Prime number")
		for k, v := range isPrime {
			if v == true {
				fmt.Printf("Prime numbers %v\n", k)
			}
		}
	}
	cntPrime := 0
	for i := 2; i < p; i++ {
		if isPrime[i] {
			cntPrime++
		}
	}
	return cntPrime
}

func main() {
	if val := os.Getenv("debugFlag"); val != "" {
		debugFlag = strings.ToLower(val) == "true"
	} else {
		debugFlag = false
	}

	getPrimeFrom := 6
	fmt.Printf("%v Prime number present in range 1-%v \n", getPrimeCount(getPrimeFrom), getPrimeFrom)
}
