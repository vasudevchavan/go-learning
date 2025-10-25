package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

var debugFlag bool = true

func findMissingElement(a []int) int {
	// Array must start from 0...N
	if a[0] != 0 {
		return 0
	} else {
		for i := 0; i < len(a)-1; i++ {
			if a[i+1]-a[i] != 1 {
				return a[i] + 1
			}
		}
	}

	// Missing element is last number
	return a[len(a)-1] + 1
}

func findMissingElementMath(a []int) int {
	n := len(a)
	expectedSum := n * (n + 1) / 2
	actualSum := 0
	for _, val := range a {
		actualSum += val
	}
	return expectedSum - actualSum
}

func findMissingElementHash(a []int) int {
	pNum := make(map[int]bool)
	for _, val := range a {
		pNum[val] = true
	}
	if debugFlag {
		for key, val := range pNum {
			fmt.Printf("%v -- %v \n", key, val)
		}
	}

	for i := 0; i < len(a)-1; i++ {
		if debugFlag {
			fmt.Printf("values inside loop is %v \n", pNum[i])
		}
		if !pNum[i] {
			return i
		}
	}

	return a[len(a)-1] + 1
}

func main() {

	if val := os.Getenv("debugFlag"); val != "" {
		debugFlag = strings.ToLower(val) == "true"
	} else {
		debugFlag = false
	}

	// Input array must be in this form 0...N
	input_Array := []int{4, 1, 3, 0}
	sorted_array := make([]int, len(input_Array))
	copy(sorted_array, input_Array)
	sort.Ints(sorted_array)
	fmt.Printf("Sorted array is %v \n", sorted_array)
	fmt.Println("\nUsing BruteForce method")
	fmt.Printf("Missing element is %v\n", findMissingElement(sorted_array))
	fmt.Println("\nUsing Mathematical guass formula")
	fmt.Printf("Missing element is %v\n", findMissingElementMath(sorted_array))
	fmt.Println("\nUsing hashMap")
	fmt.Printf("Missing element is %v\n", findMissingElementHash(sorted_array))

}
