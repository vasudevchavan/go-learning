package main

import (
	"fmt"
	"sync"
)

var (
	sendmessage string = "hey"
)



func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	fmt.Println("Refer withConcurrencyGoRoutines for goroutines Prallel execution\n")
	withConcurrencyGoRoutines()

	fmt.Println("Refer withoutConcurrencyGoRoutines for goroutines sequentially\n")
	withoutConcurrencyGoRoutines()

	fmt.Println("Refer oddEvenNumber for  Goroutines finding odd even number\n")
	oddeven := []int{2, 1, 3, 4}
	oddEvenNumber(oddeven)

	fmt.Println("Refer squareNumber for goroutines writing to a slice of integer\n")
	result := make([]int, 10)
	squareNumber(result)

	fmt.Println("go run -race . will result in Data Race error\n")
	// withOutMutex(&wg)

	fmt.Println("go run -race . and before running this comment --withOutMutex-- function \n")
	withMutex(&wg)

	stringMutex := withMutexString(8, &wg, &mu)
	fmt.Printf("String example with Mutex is %v\n", stringMutex)
	withMutexMap(8, &wg)
	rwMutexString(&wg)
	rwMutexSlice(&wg)
	rwMutexMap(&wg)

	counterWaitGroup(&wg, &mu)

	// Find total of int in a slice

	passInt := []int{1, 2, 2, 4, 7}
	fmt.Printf("Final sum of a slice is %d \n", sumOfSlice(&wg, &mu, passInt))

	// finding sum of all values in a Map

	passMap := map[string]int{
		"a": 1,
		"b": 3,
		"c": 4,
	}

	fmt.Printf("Final sum of a slice is %d \n", sumOfMap(&wg, &mu, passMap))

	//

}
