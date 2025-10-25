package main

import (
	"fmt"
)

func compareSlice(s1 []string, s2 []string) {
	cnt := 0
	for i := 0; i < len(s1); i++ {
		for j := 0; j < len(s2); j++ {
			if s1[i] == s2[j] {
				cnt++
				fmt.Printf("Matching elements arr1[%d] == arr2[%d] value:%s\n", i, j, s1[i])
			}
		}
	}
	fmt.Printf("We have found %d matching elements \n", cnt)
}

// Time Complexity == O(n^2)
// Space Complexity == O(1)

func compareSliceUsingMap(s1, s2 []string) {
	// Creata a empty Map
	cmap := make(map[string]int) // Creating a Map of n elements so space complexity is O(n)
	// store all elements from Array1 into map
	for i, v := range s1 {
		cmap[v] = i
	}
	// initialise the counter
	cnt := 0
	// matching of values
	for i, v := range s2 {
		if j, exists := cmap[v]; exists {
			cnt++
			fmt.Printf("Matching elements arr1[%d] == arr2[%d] value:%s \n", j, i, v)
		}
	}
	fmt.Printf("We have found %d matching elements\n", cnt)
}

// Time Complexity == O(n)
// Space Complexity == O(n)

func main() {
	arr1 := []string{"a", "b", "c"}
	arr2 := []string{"c", "a", "x"}

	// Function has been written using brute force
	// and this is expensive interms of time and space complexity
	fmt.Println("brute force version")
	compareSlice(arr1, arr2)

	// Second solution using hashmap
	fmt.Println("\nhashmap version")
	compareSliceUsingMap(arr1, arr2)
}
