package main

import (
	"fmt"
)

func main() {
	var arr1 [4]int
	arr1[0] = 0
	arr1[1] = 6
	arr1[2] = 5
	arr1[3] = 3
	var arr2 = [...]int{0, 1, 2, 3}
	arr3 := [4]int{0, 1, 2, 5}

	// Copying array
	arr4 := arr3
	if arr4 == arr3 {
		fmt.Println("Same array")
	}

	//
	fmt.Println("First print example")
	for i, v := range arr1 {
		fmt.Printf("value:%d-index:%d\n", i, v)
	}

	//
	fmt.Println("Second print example")
	for i := 0; i < len(arr1); i++ {
		fmt.Printf("value:%d-index:%d\n", i, arr2[i])
	}
	fmt.Println("Third print example")
	PrintArray(arr4)

	// Finding a number
	fNum := 3
	found := false
	for _, v := range arr1 {
		if v == fNum {
			found = true
			break
		}
	}
	if found {
		fmt.Printf("Value %d found \n", fNum)
	}

	// Covert Array to Slice
	sliceArr1 := arr1[:]
	for i, v := range sliceArr1 {
		fmt.Printf("Slice value:%d-index:%d\n", i, v)
	}

	// Pointer
	ptr := &arr1
	(*ptr)[0] = 99
	fmt.Printf("Printing Pointer modified array %d \n", arr1)
}

func PrintArray(a [4]int) {
	for i := 0; i < len(a); i++ {
		fmt.Printf("value:%d-index:%d\n", i, a[i])
	}
}
