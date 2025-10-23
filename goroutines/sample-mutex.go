package main

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex

// Writing to a predefined slice of integer
func writingToSlice(i int, result []int, wg *sync.WaitGroup) {
	defer wg.Done()
	result[i] = i * i
}

func squareNumber(result []int) {
	for i := 0; i < len(result); i++ {
		wg.Add(1)
		i := i
		go writingToSlice(i, result, &wg)
	}
	wg.Wait()
}

func withOutMutex(wg *sync.WaitGroup) {
	sliceOfInt := []int{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			sliceOfInt = append(sliceOfInt, num)
		}(i)
	}
	wg.Wait()
	fmt.Println("Length of numbers:", len(sliceOfInt))
}

// Sample error message we see with "withOutMutex"
// WARNING: DATA RACE
// Write at 0x00c00022e0b8 by goroutine 118:
//   main.withOutMutex.func1()
//       /Users/vasudevchavan/delete_today_junk/goroutines/sample-rwdata.go:29 +0xbc
//   main.withOutMutex.gowrap1()
//       /Users/vasudevchavan/delete_today_junk/goroutines/sample-rwdata.go:30 +0x40

// Previous write at 0x00c00022e0b8 by goroutine 116:
//   main.withOutMutex.func1()
//       /Users/vasudevchavan/delete_today_junk/goroutines/sample-rwdata.go:29 +0xbc
//   main.withOutMutex.gowrap1()
//       /Users/vasudevchavan/delete_today_junk/goroutines/sample-rwdata.go:30 +0x40

// Goroutine 118 (running) created at:
//   main.withOutMutex()
//       /Users/vasudevchavan/delete_today_junk/goroutines/sample-rwdata.go:27 +0x88
//   main.main()
//       /Users/vasudevchavan/delete_today_junk/goroutines/main.go:25 +0xa4

// Goroutine 116 (running) created at:
//   main.withOutMutex()
//       /Users/vasudevchavan/delete_today_junk/goroutines/sample-rwdata.go:27 +0x88
//   main.main()
//       /Users/vasudevchavan/delete_today_junk/goroutines/main.go:25 +0xa4

func withMutex(wg *sync.WaitGroup) {
	sliceOfInt := []int{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			mutex.Lock()
			sliceOfInt = append(sliceOfInt, num)
			mutex.Unlock()
		}(i)
	}
	wg.Wait()
	// fmt.Println("Length of numbers:", len(sliceOfInt))
}

func withMutexString(ln int, wg *sync.WaitGroup) string {
	var cstring string
	for i := 0; i < ln; i++ {
		wg.Add(1)
		go func(num int) {
			j := num
			defer wg.Done()
			mutex.Lock()
			cstring += fmt.Sprintf("%d", j)
			mutex.Unlock()
		}(i)
	}
	wg.Wait()
	return cstring
}

func withMutexMap(ln int, wg *sync.WaitGroup) {
	sharedMap := make(map[int]string)
	for i := 0; i < ln; i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			j := num
			mutex.Lock()
			sharedMap[num] = fmt.Sprintf("Value-%d", j)
			mutex.Unlock()

		}(i)
	}

	wg.Wait()

	// Print map length and some values
	fmt.Println("Map length:", len(sharedMap))
	for k, v := range sharedMap {
		fmt.Printf("%d: %s\n", k, v)
	}
}
