package main

import (
	"fmt"
	"strings"
)

func main() {
	a := true
	b := false
	b = !a
	fmt.Printf("head %v or tail %v is a == b %v", a, b, a == b)

	i := 10
	var j int = 11
	var f float32 = 11.5
	fmt.Printf("adding i %v + % v and result is %v", i, j, i+j)
	fmt.Printf("float value", f)

	var tell string = "I am in india"
	fmt.Printf("do you live in india %v", strings.HasSuffix(tell, "india"))

	var aa int = 100

	switch aa % 3 {
	case 0:
		fmt.Println("divisible by 3")
	default:
		fmt.Println("not divisible")
	}

	// Looping using index

	for _, i := range []int{10, 11, 12, 13, 14, 15, 16, 17, 18} {
		fmt.Println(i)
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	ab := 2
	for ab < 10 {
		fmt.Println(ab)
		ab = ab + 2
	}

	
}
