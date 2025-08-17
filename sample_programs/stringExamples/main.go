package main

import (
	"fmt"
	"strings"
)

func main() {
	// isDebug := false
	str := "oneonetwoonnoe two"
	fString := "two"
	// Using inbuilt function we are counting
	// number of matching substring
	numberOfMatching(str, fString)
	matchingSubstring(str, fString)

}

func numberOfMatching(str string, fString string) {
	n := 0
	for {
		i := strings.Index(str, fString)
		if i == -1 {
			break
		}
		n++
		str = str[i+len(fString):]
	}
	fmt.Printf("We have found %d occurrences of sub strings %s \n", n, fString)
}

func matchingSubstring(str string, fString string) {
	isDebug := false
	cnt := 0
	capture_Index := []int{}
	lStr := len(str)
	lSubString := len(fString)
	for i := 0; i <= lStr-lSubString; i++ {
		if isDebug {
			fmt.Printf("Length of whole string %d and current posting %d \n", lStr, i)
		}
		match := true
		for j := 0; j < lSubString; j++ {
			if isDebug {
				fmt.Printf("String matching testing on %s --- %s \n", string(str[i+j]), string(fString[j]))
			}
			if str[i+j] != fString[j] {
				match = false
				break
			}
		}

		if match {
			cnt++
			// Capturing starting position of the index
			// so that this can be used to replace matching string at nth position.
			capture_Index = append(capture_Index, i)
		}
	}

	fmt.Printf("We have found %d occurrences of sub strings %s \n", cnt, fString)
	fmt.Printf("Starting index of Substring %d \n", capture_Index)
}
