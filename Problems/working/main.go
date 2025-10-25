package main

import "fmt"

func main() {
	givenIntArraySlice := []int{1, 2, 1, 2, 3, 4, 4, 3, 5}
	givenIntArrayMap := make(map[int]int)

	// convert slice of int to Map
	for _, num := range givenIntArraySlice {
		givenIntArrayMap[num]++
	}
	for k, v := range givenIntArrayMap {
		fmt.Println(k, v)
	}

	// convert slice of string to Map
	givenStringArraySlice := []string{"apple", "apricot", "banana", "butterfruit", "chikoo", "cheery", "mango"}
	givenStringArrayMap := make(map[string][]string)

	for _, fruit := range givenStringArraySlice {
		firstLetter := string(fruit[0])
		givenStringArrayMap[firstLetter] = append(givenStringArrayMap[firstLetter], fruit)
	}
}
