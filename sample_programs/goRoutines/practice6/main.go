package main

import (
	"fmt"
	"time"
)

func echoo(ping chan string, pong chan string) {
	for {
		s := <-ping
		pong <- fmt.Sprintf("channel have received %s and sending it back", s)
	}
}

func main() {
	ping := make(chan string)
	pong := make(chan string)

	go echoo(ping, pong)

	checking := []string{"hello", "world"}

	for _, v := range checking {
		ping <- v
		msgPong := <-pong
		fmt.Println(msgPong)
	}
	close(ping) // Stop the Go Routines
	time.Sleep(1 * time.Second)
}
