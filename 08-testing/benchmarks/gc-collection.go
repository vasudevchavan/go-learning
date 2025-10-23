package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"runtime/debug"
)

func printStats(mem runtime.MemStats) {
	runtime.ReadMemStats(&mem)
	fmt.Println("mem.Alloc", mem.Alloc)
	fmt.Println("mem.TotalAlloc", mem.TotalAlloc)
	fmt.Println("mem.HeapAlloc", mem.HeapAlloc)
	fmt.Println("mem.NumGC", mem.NumGC)
}

func main() {
	go func() {
		fmt.Println("Starting pprof server on http://localhost:8081/debug/pprof/")
		err := http.ListenAndServe(":8081", nil)
		if err != nil {
			fmt.Println("pprof server error:", err)
		}
	}()

	var mem runtime.MemStats
	var s []byte
	printStats(mem)

	for i := 0; i < 10; i++ {
		s = make([]byte, 10_000_000_000)
		if s == nil {
			fmt.Println("Ops failed")
		}
	}

	s = nil
	runtime.GC()
	debug.FreeOSMemory()
	printStats(mem)

	select {} // Keeps the server alive for pprof

	// Command to run this program
	// GODEBUG=gctrace=1 go run gcollection.go
}
