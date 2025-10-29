package main

import (
	"fmt"
)

func main() {
	t := testClosure()
	fmt.Println(t())
	fmt.Println(t())
	fmt.Println(t())
	fmt.Println(t())
	fmt.Println(t())
	fmt.Println(t())
}

func testClosure() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
