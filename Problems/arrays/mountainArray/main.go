package main

import "fmt"

func findMountainArray(arr []int) {
	i := 1
	for i < len(arr) && arr[i] > arr[i-1] {
		i++
	}
	if i == 1 || i == len(arr) {
		fmt.Println("Not a mountain array")
	}

	for i < len(arr) && arr[i] < arr[i-1] {
		i++
	}
	if i == len(arr) {
		fmt.Println("It's a mountain array")
	} else {
		fmt.Println("Not a mountain array")
	}
}

func main() {
	array_int := []int{1, 2, 3, 1, 4}

	if len(array_int) < 1 {
		fmt.Println("Invalid array")
	} else {
		findMountainArray(array_int)
	}
}
