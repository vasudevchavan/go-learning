package main

import (
	"fmt"
	// "sync"
	"context"
	"time"
)



func doWork(ctx context.Context){

	for{
		select {
		case <-ctx.Done():
			fmt.Println("GoRoutine: Received cancellatin")
		default:
			fmt.Println("GoRoutine: Working")
			time.Sleep(500 * time.Second)
		}
	}

}


func main() {

	ctx,cancel := context.WithCancel(context.Background())

	go doWork(ctx)

	time.Sleep(2 * time.Second)

	fmt.Println("Main: cancelling all workers")

	cancel()

	time.Sleep(1 * time.Second)
	fmt.Println("Main: done")
}