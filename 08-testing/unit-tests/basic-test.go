package main

import (
	"fmt"
	"sync"
)

var (
	mu sync.Mutex
	wg sync.WaitGroup
)

func counter_waitGroup(wg *sync.WaitGroup, mu *sync.Mutex) {
	var count int
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			count++
			mu.Unlock()
		}()

	}
	wg.Wait()
	fmt.Printf("value of final count %d \n", count)
}

func sumOfSlice(wg *sync.WaitGroup, mu *sync.Mutex, arr []int) (finalSum int) {
	var sum int = 0
	for _, v := range arr {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			mu.Lock()
			sum += num
			mu.Unlock()
		}(v)
	}
	wg.Wait()

	return sum

}

func sumOfMap(wg *sync.WaitGroup, mu *sync.Mutex, m map[string]int) (finalSum int) {
	var sum int = 0

	for _, v := range m {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			mu.Lock()
			sum += num
			mu.Unlock()
		}(v)
	}
	wg.Wait()
	return sum

}

func main() {
	
	counter_waitGroup(&wg, &mu)
	
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






