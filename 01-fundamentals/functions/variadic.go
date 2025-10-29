package main

import (
	"fmt"
)

func main() {

	checking("start")
	total := sumofnum(1, 2, 3, 4, 5, 6, 7)
	fmt.Printf("Sum of numbers is %d\n", total)

	// defer function whose execution is deferred
	// to the moment the surrounding function returns
	defer checking("middle")

	getNum := []int{1, 2, 3, 4, 5, 6, 7}
	nTotal := sumofnum(getNum...)
	fmt.Printf("Sum of numbers is %d", nTotal)

	checking("last")

}

func sumofnum(nn ...int) int {
	fmt.Printf("Input of type:%T\n", nn)
	t := 0
	for _, v := range nn {
		t += v
	}
	return t
}

func checking(s string) {
	fmt.Println(s, "\n")
}
