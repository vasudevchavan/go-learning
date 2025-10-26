package utils

import (
	"container/heap"
	"fmt"
	"sort"
)

// -------------------------
// 🔁 Reverse Map
// -------------------------

// ReverseMap swaps keys and values in a map.
// Note: If values are not unique, later keys overwrite earlier ones.
func ReverseMap[K comparable, V comparable](m map[K]V) map[V]K {
	reversed := make(map[V]K)
	for k, v := range m {
		reversed[v] = k
	}
	return reversed
}

// -------------------------
// 🧮 Top-K Frequent Elements
// -------------------------

type freqItem struct {
	value string
	count int
}

type freqHeap []freqItem

func (h freqHeap) Len() int           { return len(h) }
func (h freqHeap) Less(i, j int) bool { return h[i].count > h[j].count } // Max-heap
func (h freqHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *freqHeap) Push(x any)        { *h = append(*h, x.(freqItem)) }
func (h *freqHeap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

// TopKFrequent finds the top K most frequent words in a slice.
func TopKFrequent(words []string, k int) []string {
	count := make(map[string]int)
	for _, w := range words {
		count[w]++
	}

	h := &freqHeap{}
	for word, freq := range count {
		heap.Push(h, freqItem{word, freq})
	}

	result := []string{}
	for i := 0; i < k && h.Len() > 0; i++ {
		item := heap.Pop(h).(freqItem)
		result = append(result, item.value)
	}
	return result
}

// -------------------------
// 🔡 Group Anagrams
// -------------------------

// GroupAnagrams groups words that are anagrams of each other.
func GroupAnagrams(words []string) map[string][]string {
	grouped := make(map[string][]string)
	for _, w := range words {
		runes := []rune(w)
		sort.Slice(runes, func(i, j int) bool { return runes[i] < runes[j] })
		key := string(runes)
		grouped[key] = append(grouped[key], w)
	}
	return grouped
}

// -------------------------
// ➕ Two Sum
// -------------------------

// TwoSum returns the indices of two numbers that add up to target.
func TwoSum(nums []int, target int) [2]int {
	m := make(map[int]int)
	for i, n := range nums {
		if j, ok := m[target-n]; ok {
			return [2]int{j, i}
		}
		m[n] = i
	}
	return [2]int{-1, -1} // Not found
}

// -------------------------
// 🔄 Merge Nested Maps
// -------------------------

// MergeNestedMaps merges two maps[string]map[string]int into one.
func MergeNestedMaps(m1, m2 map[string]map[string]int) map[string]map[string]int {
	result := make(map[string]map[string]int)

	// Deep copy m1
	for k, v := range m1 {
		result[k] = make(map[string]int)
		for innerK, innerV := range v {
			result[k][innerK] = innerV
		}
	}

	// Merge m2 into result
	for k, v := range m2 {
		if result[k] == nil {
			result[k] = make(map[string]int)
		}
		for innerK, innerV := range v {
			result[k][innerK] = innerV
		}
	}
	return result
}

// -------------------------
// 🧪 Test Runner for Advanced Map Operations
// -------------------------

func TestAdvancedMapOperations() {
	fmt.Println("\n🧠 Advanced Map Algorithms")

	// 1️⃣ Reverse Map
	m1 := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println("Original:", m1)
	fmt.Println("Reversed:", ReverseMap(m1))

	// 2️⃣ Top K Frequent
	words := []string{"go", "is", "fun", "go", "is", "fast", "go", "go", "fun"}
	fmt.Println("\nTop 2 Frequent Words:", TopKFrequent(words, 2))

	// 3️⃣ Group Anagrams
	wordList := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	grouped := GroupAnagrams(wordList)
	fmt.Println("\nGrouped Anagrams:")
	for k, v := range grouped {
		fmt.Printf("Key: %-10s -> %v\n", k, v)
	}

	// 4️⃣ Two Sum
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Printf("\nTwoSum(%v, %d) = %v\n", nums, target, TwoSum(nums, target))

	// 5️⃣ Merge Nested Maps
	m2 := map[string]map[string]int{
		"Suresh": {"Math": 80, "English": 70},
	}
	m3 := map[string]map[string]int{
		"Suresh": {"Social": 90},
		"Ramesh": {"Arts": 100},
	}
	merged := MergeNestedMaps(m2, m3)
	fmt.Println("\nMerged Nested Maps:", merged)
}

