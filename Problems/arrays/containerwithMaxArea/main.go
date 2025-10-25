package main

import "fmt"

func getContainerArea(x int, y int, z int) int {
	var area = 0
	if x > y {
		area = y * z
	} else {
		area = x * z
	}
	return area
}

func finalContainerArea(arr []int) (int, int, int) {
	f_area, g_sum, s_index, e_index := 0, 0, 0, 0
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			g_sum = getContainerArea(arr[i], arr[j], j-i)
			if g_sum > f_area {
				f_area = g_sum
				s_index = i
				e_index = j
			}
		}
	}
	// fmt.Println(f_area, s_index, e_index)
	return s_index, e_index, f_area
}

func finalContainerArea_optimized(arr []int) (int, int, int) {
	left_index, s_left_index, right_index, s_right_index, maxArea, area, height, width := 0, 0, 0, 0, 0, 0, 0, 0
	left_index = 0
	right_index = len(arr) - 1
	for left_index < right_index {
		width = right_index - left_index
		if arr[left_index] > arr[right_index] {
			height = arr[right_index]
			area = height * width
			if area > maxArea {
				maxArea, s_left_index, s_right_index = area, left_index, right_index
			}
			right_index--
		} else {
			height = arr[left_index]
			area = height * width
			if area > maxArea {
				maxArea, s_left_index, s_right_index = area, left_index, right_index
			}
			left_index++
		}

	}
	// fmt.Println(s_left_index, s_right_index, maxArea)
	return s_left_index, s_right_index, maxArea
}

func getValue(a1 int, b1 int, cal int) int {
	getvalue := 0
	if cal == 0 {
		if a1 > b1 {
			getvalue = b1
		} else {
			getvalue = a1
		}
	} else if cal == 1 {
		if a1 > b1 {
			getvalue = a1
		} else {
			getvalue = b1
		}
	}
	return getvalue
}

func brututeForceApproach(arr []int) int {
	max_area, area, wid, leng := 0, 0, 0, 0
	n := len(arr)

	for l1 := 0; l1 < n; l1++ {
		for l2 := l1 + 1; l2 < n; l2++ {
			leng = getValue(arr[l1], arr[l2], 0)
			wid = l2 - l1
			area = leng * wid
			if area > max_area {
				max_area = area
			}
		}
	}

	return max_area
}

func twoPointerArea(arr []int) int {
	max_area, area, wid, leftPtr, rightPtr, leg := 0, 0, 0, 0, 0, 0

	leftPtr = 0
	rightPtr = len(arr) - 1

	for leftPtr < rightPtr {
		leg = min(arr[leftPtr], arr[rightPtr])
		wid = rightPtr - leftPtr
		area = leg * wid
		max_area = max(max_area, area)
		if arr[leftPtr] < arr[rightPtr] {
			leftPtr++
		} else {
			rightPtr--
		}
	}
	return max_area
}

func main() {
	arrayInt := []int{1, 2000, 1, 20, 1000, 200, 10, 10, 999}
	finalContainerArea(arrayInt)
	finalContainerArea_optimized(arrayInt)

	// Brute force approach
	maxarea := brututeForceApproach(arrayInt)
	fmt.Println("Max area is: ", maxarea)

	// Two Pointer approach
	maxby2ptr := twoPointerArea(arrayInt)
	fmt.Println("Max area by 2 pointer is: ", maxby2ptr)

}
