package main

import (
	"fmt"
	"os"
	"strings"
)

var debugFlag bool = false

// Use this instead of math.Max function which does some additonal conversion
func max_value(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestSubString(s string) int {
	maxLength, leftPointer := 0, 0
	// rightPointer := 0
	validMap := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if val, ok := validMap[s[i]]; ok {
			// leftPointer = int(math.Max(float64(leftPointer), float64(val+1)))
			leftPointer = max_value(leftPointer, val+1)
		}
		// rightPointer += 1
		validMap[s[i]] = i
		// maxLength = int(math.Max(float64(maxLength), float64(i-leftPointer+1)))
		maxLength = max_value(maxLength, i-leftPointer+1)
		if debugFlag {
			fmt.Printf("Left pointer: %v, Right pointer: %v, Max length: %v \n", leftPointer, i, maxLength)
			for key, value := range validMap {
				fmt.Printf("Key: %c, Value: %v \n", key, value)
			}
		}
	}
	return maxLength
}

func main() {
	if val := os.Getenv("debugFlag"); val != "" {
		debugFlag = strings.ToLower(val) == "true"
	} else {
		debugFlag = false
	}
	givenString := "abaaa"
	fmt.Printf("Longest substring in the given string --%v is %v \n", givenString, longestSubString(givenString))
}
