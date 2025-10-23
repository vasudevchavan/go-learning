package main

import (
	"fmt"
)

func printSomething(msg string) {
	fmt.Println("Hello world")
}

func main() {
	var sendMessage string = "Hello world"
	printSomething(sendMessage)
	sendMessage = "Hello home"
	printSomething(sendMessage)
}
