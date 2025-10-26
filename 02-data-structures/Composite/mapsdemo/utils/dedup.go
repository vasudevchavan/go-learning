package utils

// RemoveDuplicateArray removes duplicates and returns unique items + counts.
func RemoveDuplicateArray(arr []string) ([]string, map[string]int) {
	uniqueMap := make(map[string]bool)
	countMap := make(map[string]int)

	for _, v := range arr {
		uniqueMap[v] = true
		countMap[v]++
	}

	uniqueArr := make([]string, 0, len(uniqueMap))
	for k := range uniqueMap {
		uniqueArr = append(uniqueArr, k)
	}

	return uniqueArr, countMap
}

// RemoveDuplicateChars counts characters and returns frequency map.
func RemoveDuplicateChars(s string) map[rune]int {
	charCount := make(map[rune]int)
	for _, ch := range s {
		charCount[ch]++
	}
	return charCount
}


// Reverse Map ignoring duplicate
func ReverseMapIgnoringDup(samplemap map[string]int) (rev map[int]string){
	reverseMap := make(map[int]string)
	for k, v := range samplemap {
		reverseMap[v] = k
	}
	return reverseMap
}

// Reverse Map considering duplicate
func ReverseMapConsiderDup(samplemap map[string]int)(rev map[int][]string){
	reverseMapDup := make(map[int][]string)
	for k, v := range samplemap {
		reverseMapDup[v] = append(reverseMapDup[v],k)
	}
	return reverseMapDup
}

