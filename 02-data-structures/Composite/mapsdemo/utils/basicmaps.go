package utils

import (
	"fmt"
	"sort"
)

// BasicMapOperations demonstrates creation, update, delete, lookup
func BasicMapOperations() {
	fmt.Println("1️⃣ Basic Map Operations")
	rollnum := map[string]int{
		"C":   1,
		"C++": 2,
		"Java": 3,
	}
	rollnum["Go"] = 4
	delete(rollnum, "Java")
	rollnum["Java"] = 3

	if v, exists := rollnum["C"]; exists {
		fmt.Printf("Found C -> %d\n", v)
	}

	PrintMap(rollnum)
	fmt.Printf("Map length: %d\n\n", len(rollnum))
}

// ArrayToMapConversion shows how to convert arrays to maps
func ArrayToMapConversion() {
	fmt.Println("2️⃣ Array to Map Conversion")
	langs := [4]string{"Go", "C", "C++", "C#"}
	langMap := make(map[int]string)
	for i, v := range langs {
		langMap[i] = v
	}
	PrintMap(langMap)

	codes := [4]int{56001, 56002, 56008, 56098}
	areas := [4]string{"BSK", "Jayanagar", "JP Nagar", "RR Nagar"}
	postalMap := make(map[string]int)
	for i := 0; i < len(codes); i++ {
		postalMap[areas[i]] = codes[i]
	}
	PrintMap(postalMap)
	fmt.Println()
}

// DuplicatesHandling shows removing duplicates and counting
func DuplicatesHandling() {
	fmt.Println("3️⃣ Duplicates Handling")
	fruits := [6]string{"apple", "apple", "mango", "orange", "mango", "kiwi"}

	uniqueFruits := make(map[string]bool)
	for _, f := range fruits {
		uniqueFruits[f] = true
	}
	PrintMap(uniqueFruits)
	PrintMapAsJSON(uniqueFruits)

	fruitCount := make(map[string]int)
	for _, f := range fruits {
		fruitCount[f]++
	}
	fmt.Println("Duplicate count:")
	for k, v := range fruitCount {
		if v > 1 {
			fmt.Printf("%s has %d duplicates\n", k, v)
		}
	}
	fmt.Println()
}

// MapsOfSlices shows how to use slices as map values
func MapsOfSlices() {
	fmt.Println("4️⃣ Maps of Slices")
	friends := map[string][]string{
		"Alice": {"Bob", "Charlie"},
		"Bob":   {"Alice"},
	}
	for name, flist := range friends {
		fmt.Printf("%s has friends: %v\n", name, flist)
	}
	fmt.Println()
}

// NestedMaps demonstrates nested map usage
func NestedMaps() {
	fmt.Println("5️⃣ Nested Maps")
	users := map[string]map[string]string{
		"alice": {"email": "alice@mail.com", "city": "London"},
		"bob":   {"email": "bob@mail.com", "city": "New York"},
	}
	for user, info := range users {
		fmt.Printf("%s -> %v\n", user, info)
	}
	fmt.Println("Alice's email:", users["alice"]["email"])
	fmt.Println()
}

// StructAsMapKey demonstrates structs as map keys
func StructAsMapKey() {
	fmt.Println("6️⃣ Struct as Map Key")
	type Point struct{ X, Y int }
	points := map[Point]string{
		{1, 2}:  "Top Right",
		{-1, -2}: "Bottom Left",
	}
	PrintMap(points)
	fmt.Println()
}

// SortedKeyIteration shows sorted map iteration
func SortedKeyIteration() {
	fmt.Println("7️⃣ Sorted Map Iteration")
	scores := map[string]int{"Alice": 95, "Bob": 87, "Eve": 92}
	keys := make([]string, 0, len(scores))
	for k := range scores {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("%s -> %d\n", k, scores[k])
	}
}

