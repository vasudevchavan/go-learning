package main

import (
	"fmt"
	"os"
	"strings"
)

var debugFlag = true

// time complexity:
//  O(n) only one for loop to read all elements of the array
// space complexity
// O(1) since no extra memory is used and valus are just swapped

func moveZeroes(nums []int) []int {
	for i, j := 0, 0; i < len(nums); i++ {
		if debugFlag {
			fmt.Println("i", i, "j", j, "num of i ", nums[i], "num of j", nums[j])
		}
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
	return nums
}

// time complexity:
//	O(n) for first loop and O(k) is for the second loop
// since k < n time complexity is O(n)
// space complexity
// we are creating a new empty slice
// which hold the same n number of elements & its O(n)

func secondMoveZeroes(num2 []int) []int {
	letsSortSlice := []int{}
	cntZero := 0
	for i := 0; i < len(num2); i++ {
		if num2[i] == 0 {
			cntZero++
		} else {
			letsSortSlice = append(letsSortSlice, num2[i])
		}
	}

	for i := 0; i < cntZero; i++ {
		letsSortSlice = append(letsSortSlice, 0)
	}

	return letsSortSlice
}

func main() {

	if val := os.Getenv("debugFlag"); val != "" {
		debugFlag = strings.ToLower(val) == "true"
	} else {
		debugFlag = false
	}
	slice_int := []int{0, 1, 0, 3, 12}
	fmt.Printf("Final sorted array is %v\n", moveZeroes(slice_int))
	fmt.Printf("Second method to sort 0's to the last %v\n", secondMoveZeroes(slice_int))

}
