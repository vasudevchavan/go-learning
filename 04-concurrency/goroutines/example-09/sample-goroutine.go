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

	for i := 0; i < 10; i++ {
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

type SafeString struct {
	murw sync.RWMutex
	str  string
}

type SafeSlice struct {
	murw  sync.RWMutex
	slice []int
}

type SafeMap struct {
	murw sync.RWMutex
	m    map[string]int
}

func (sm *SafeMap) Read(key string) (int, bool) {
	sm.murw.RLock()
	defer sm.murw.RUnlock()
	val, ok := sm.m[key]
	return val, ok
}

func (sm *SafeMap) Write(key string, value int) {
	sm.murw.Lock()
	defer sm.murw.Unlock()
	sm.m[key] = value
}

func (ss *SafeSlice) Read(index int) (int, bool) {
	ss.murw.RLock()
	defer ss.murw.RUnlock()
	if index < 0 || index >= len(ss.slice) {
		return 0, false
	}
	return ss.slice[index], true

}

func (ss *SafeSlice) Write(index int, value int) {
	ss.murw.Lock()
	defer ss.murw.Unlock()
	if index >= len(ss.slice) {
		newSlice := make([]int, index+1)
		copy(newSlice, ss.slice)
		ss.slice = newSlice
	}
	ss.slice[index] = value
	fmt.Printf("Written %d at index %d\n", value, index)
}

func (ss *SafeString) Read() string {
	// One way
	// ss.mu.RLock()
	// defer ss.mu.RUnlock()
	// return ss.str

	// Second way
	ss.murw.RLock()
	val := ss.str // Copy value
	defer ss.murw.RUnlock()
	return val
}

func (ss *SafeString) Write(newStr string) {
	ss.murw.Lock()
	defer ss.murw.Unlock()
	ss.str = newStr
	fmt.Println("Update string to:", newStr)
}

func rwMutexString(wg *sync.WaitGroup) {
	ss := SafeString{str: "String"}
	// Writing to string
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			ss.Write(fmt.Sprintf("string-%d", i))
		}(i)
	}
	// Reading from string
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			val := ss.Read()
			fmt.Printf("Reader %d read String %v \n", id, val)
		}(i)

	}
	wg.Wait() // This defines that this program runs in parallel
}

func rwMutexSlice(wg *sync.WaitGroup) {
	ss := SafeSlice{slice: make([]int, 0)}
	// Writers: write values concurrently
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			ss.Write(i, i*10)
		}(i)
	}
	// Readers: read values concurrently
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			val, ok := ss.Read(i)
			if ok {
				fmt.Printf("Read index %d: %d\n", i, val)
			} else {
				fmt.Printf("Read index %d: out of range\n", i)
			}
		}(i)
	}
	wg.Wait()
	fmt.Println("Final slice:", ss.slice)
}

func rwMutexMap(wg *sync.WaitGroup) {
	sm := &SafeMap{m: make(map[string]int)}
	// Writers: write values concurrently
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := fmt.Sprintf("key-%d", i)
			sm.Write(key, i*100)
		}(i)
	}
	// Readers: read values concurrently
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := fmt.Sprintf("key-%d", i)
			val, ok := sm.Read(key)
			if ok {
				fmt.Printf("Reader %d read: key=%s, value=%d\n", i, key, val)
			} else {
				fmt.Printf("Reader %d read: key=%s not found\n", i, key)
			}
		}(i)
	}
	wg.Wait()

	for key, val := range sm.m {
		fmt.Printf(" key:%v,val:%v \n", key, val)
	}
}

// Writing to a predefined slice of integer
func writingToSlice(i int, result []int, wg *sync.WaitGroup) {

	defer wg.Done()
	result[i] = i * i
}

func squareNumber(result []int) {
	var wg sync.WaitGroup

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
	var mu sync.Mutex
	sliceOfInt := []int{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			mu.Lock()
			sliceOfInt = append(sliceOfInt, num)
			mu.Unlock()
		}(i)
	}
	wg.Wait()
	// fmt.Println("Length of numbers:", len(sliceOfInt))
}

func withMutexString(ln int, wg *sync.WaitGroup, mu *sync.Mutex) string {

	var cstring string
	for i := 0; i < ln; i++ {
		wg.Add(1)
		go func(num int) {
			j := num
			defer wg.Done()
			mu.Lock()
			cstring += fmt.Sprintf("%d", j)
			mu.Unlock()
		}(i)
	}
	wg.Wait()
	return cstring
}

func withMutexMap(ln int, wg *sync.WaitGroup) {

	var mu sync.Mutex
	sharedMap := make(map[int]string)
	for i := 0; i < ln; i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			j := num
			mu.Lock()
			sharedMap[num] = fmt.Sprintf("Value-%d", j)
			mu.Unlock()

		}(i)
	}

	wg.Wait()

	// Print map length and some values
	fmt.Println("Map length:", len(sharedMap))
	for k, v := range sharedMap {
		fmt.Printf("%d: %s\n", k, v)
	}
}

func counterWaitGroup(wg *sync.WaitGroup, mu *sync.Mutex) {

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
	sum := 0

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
