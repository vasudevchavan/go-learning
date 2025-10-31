package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	cnt := 0
	var wg sync.WaitGroup
	var RW sync.RWMutex
	fmt.Printf("CPU's:%d\n", runtime.NumCPU())

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			RW.Lock()
			cnt++
			defer RW.Unlock()
		}()
	}
	wg.Wait()

	fmt.Printf("values of count=%d\n", cnt)


  fmt.Println("Increment a Counter without Locking")
	var cnt11 int64
	// cnt111:=0
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&cnt11,1)
		}()
	}
	wg.Wait()
	fmt.Printf("values of count=%d", cnt11)
}
