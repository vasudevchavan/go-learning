package main

import (
	"fmt"
)

func test1() {
	fmt.Println("function without Return")
}

func test2(s string) {
	fmt.Printf("passing value:%s to function\n", s)
}

func test3(s string) string {
	return fmt.Sprintf("%s", s)
}

func main() {
	test1()
	test2("hey")
	fmt.Printf("function returing a value:%s\n", test3("gophers"))
}

func Add(a int, b int) int {
	return a + b
}
