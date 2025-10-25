package main

import "fmt"

func findUnpairElement(a []int) int {
	aMap := make(map[int]int) // O(1)  time complexity | O(n) space complexity
	for _, v := range a {
		aMap[v]++ // O(1) per insert and total of O(n) time complexity
	}
	for k, v := range aMap {
		if v == 1 {
			return k // O(k) time incase of k unique numbers
		}
	}
	return -1
}

// Total time complexity is O(n)
// Total space complexity is O(n)

func findUnpairElementXor(a []int) int {
	result := 0
	for _, v := range a {
		result ^= v // O(n) time complexity | O(1) for 1 unpaired element
	}
	return result
}

func main() {
	givenIntArraySlice := []int{1, 2, 1}
	fmt.Printf("Give Pair of Slice is %v\n", givenIntArraySlice)
	fmt.Println("BruteForce Method")
	fmt.Printf("Missing single pair is %v\n", findUnpairElement(givenIntArraySlice))
	fmt.Println("XOR Method")
	fmt.Printf("Missing single pair is %v\n", findUnpairElementXor(givenIntArraySlice))

}
