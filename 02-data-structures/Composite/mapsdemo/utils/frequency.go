package utils

// FrequencyCounter counts occurrences of strings in a slice.
func FrequencyCounter(items []string) map[string]int {
	counter := make(map[string]int)
	for _, v := range items {
		counter[v]++
	}
	return counter
}

