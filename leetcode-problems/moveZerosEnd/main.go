package main

import "fmt"

func moveZeroes(nums []int) []int {
	for i, j := 0, 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
	return nums
}

func main() {
	array_int := []int{0, 1, 0, 3, 12}

	fmt.Printf("Final sorted array is %v\n", moveZeroes(array_int))
}
