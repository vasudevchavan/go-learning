package main

import (
	"fmt"
	"sync"
)

type SafeString struct {
	mu  sync.RWMutex
	str string
}

type SafeSlice struct {
	mu    sync.RWMutex
	slice []int
}

type SafeMap struct {
	mu sync.RWMutex
	m  map[string]int
}

func (sm *SafeMap) Read(key string) (int, bool) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	val, ok := sm.m[key]
	return val, ok
}

func (sm *SafeMap) Write(key string, value int) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.m[key] = value
}

func (ss *SafeSlice) Read(index int) (int, bool) {
	ss.mu.RLock()
	defer ss.mu.RUnlock()
	if index < 0 || index >= len(ss.slice) {
		return 0, false
	}
	return ss.slice[index], true

}

func (ss *SafeSlice) Write(index int, value int) {
	ss.mu.Lock()
	defer ss.mu.Unlock()
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
	ss.mu.RLock()
	val := ss.str // Copy value
	defer ss.mu.RUnlock()
	return val
}

func (ss *SafeString) Write(newStr string) {
	ss.mu.Lock()
	defer ss.mu.Unlock()
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
