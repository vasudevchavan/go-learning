package main

import (
	"fmt"
)

func main() {

	fmt.Println("CallBack function example")
	fmt.Println(doMath(1, 1, add))
	fmt.Println(doMath(1, 1, sub))
	fmt.Println(doMath(1, 1, mul))
}

func add(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func mul(a int, b int) int {
	return a * b
}
func doMath(a int, b int, f func(int, int) int) int {
	return f(a, b)
}
