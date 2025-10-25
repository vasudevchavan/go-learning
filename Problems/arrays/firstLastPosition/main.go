package main

import (
	"fmt"
	"sort"
)

func firstLastPositionBruteForce(a []int, t int) (int, int) {
	first, last := -1, -1
	for i := 0; i < len(a); i++ {
		if a[i] == t {
			first = i
			break
		}
	}
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] == t {
			last = i
			break
		}
	}
	return first, last
}

func firtNumBinarySearch(a []int, t int) int {
	left, right := 0, len(a)-1
	for right >= left {
		mid := (left + right) / 2
		if a[mid] == t {
			if mid == 0 || a[mid-1] != t {
				return mid
			}
			right = mid - 1
		} else if a[mid] > t {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func lastNumberBinarySearch(a []int, t int) int {
	left, right := 0, len(a)-1
	for right >= left {
		mid := (left + right) / 2
		if a[mid] == t {
			if mid == len(a)-1 || a[mid+1] != t {
				return mid
			}
			left = mid + 1
		} else if a[mid] > t {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func main() {
	givenArray := []int{1, 1, 5, 6, 7, 3, 8, 2, 2, 3, 5, 6}
	sortArray := make([]int, len(givenArray))
	copy(sortArray, givenArray)
	sort.Ints(sortArray)
	target := 1
	first, last := firstLastPositionBruteForce(sortArray, target)
	fmt.Printf("Search first and last occurance of number in an array %v for target %v \n", sortArray, target)
	fmt.Printf("Using BruteForce method to search first and last occurance of number in an array \n")
	fmt.Printf("First occurance at %v \nLast occurance at %v \n", first, last)

	fmt.Printf("Using BinarySearch  method to search first and last occurance of number \n")
	fmt.Printf("First occurance at %v \nLast occurance at %v \n", firtNumBinarySearch(sortArray, target), lastNumberBinarySearch(sortArray, target))

}
