package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"
)

// -------------------------
// Generic Utility Functions
// -------------------------

// PrintMap prints key-value pairs of any map
func PrintMap[K comparable, V any](m map[K]V) {
	fmt.Println("Printing map:")
	for k, v := range m {
		fmt.Printf("Key: %-10v | Value: %v\n", k, v)
	}
}

// PrintMapAsJSON prints a map as pretty JSON
func PrintMapAsJSON[K comparable, V any](m map[K]V) {
	data, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		log.Println("Error marshalling map to JSON:", err)
		return
	}
	fmt.Println("Map as JSON:\n", string(data))
}



// Map with key,val as string
// SimplePrintMap unformatted
func simplePrintMap(m map[string]string) {
	fmt.Println("Map contents:", m)
}

// PrintMap formatted with key,val as string
func printMapKeyValue(m map[string]string) {
	for key, value := range m {
		fmt.Printf("Key: %-6s, | Value: %s\n", key, value)
	}
}


func printAsJson(m map[string]string) {
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		log.Println("Error marshalling map to JSON:", err)
	}
	fmt.Println("Map as JSON:", string(jsonData))
}

// -------------------------
// Main Program
// -------------------------

func main() {
	// -------------------------
	// 1. Basic Map Operations
	// -------------------------
	fmt.Println("1️⃣ Basic Map Operations")
	rollnum := make(map[string]int)
	rollnum["C"] = 1
	rollnum["C++"] = 2
	rollnum["Java"] = 3

	// Add/update/delete
	rollnum["Go"] = 4
	delete(rollnum, "Java")
	rollnum["Java"] = 3

	// Lookup with comma-ok
	if v, exists := rollnum["C"]; exists {
		fmt.Printf("Found C -> %d\n", v)
	}

	PrintMap(rollnum)
	fmt.Printf("Map length: %d\n\n", len(rollnum))

	// -------------------------
	// 2. Array -> Map Conversion
	// -------------------------
	fmt.Println("2️⃣ Array to Map Conversion")
	langs := [4]string{"Go", "C", "C++", "C#"}
	langMap := make(map[int]string)
	for i, v := range langs {
		langMap[i] = v
	}
	PrintMap(langMap)

	// Two arrays -> one map
	codes := [4]int{56001, 56002, 56008, 56098}
	areas := [4]string{"BSK", "Jayanagar", "JP Nagar", "RR Nagar"}
	postalMap := make(map[string]int)
	for i := 0; i < len(codes); i++ {
		postalMap[areas[i]] = codes[i]
	}
	PrintMap(postalMap)
	fmt.Println()

	// -------------------------
	// 3. Removing Duplicates / Counting
	// -------------------------
	fmt.Println("3️⃣ Duplicates Handling")
	fruits := [6]string{"apple", "apple", "mango", "orange", "mango", "kiwi"}

	// Remove duplicates using map[string]bool
	uniqueFruits := make(map[string]bool)
	for _, f := range fruits {
		uniqueFruits[f] = true
	}
	PrintMap(uniqueFruits)
	PrintMapAsJSON(uniqueFruits)

	// Count duplicates
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

	// -------------------------
	// 4. Maps of Slices
	// -------------------------
	fmt.Println("4️⃣ Maps of Slices")
	friends := map[string][]string{
		"Alice": {"Bob", "Charlie"},
		"Bob":   {"Alice"},
	}
	for name, flist := range friends {
		fmt.Printf("%s has friends: %v\n", name, flist)
	}
	fmt.Println()


	studentsScore := map[string]map[string]int{
		"Suresh": {
			"Math":    80,
			"English": 70,
		},
		"Anand": {
			"Math":    90,
			"English": 80,
		},
	}
	studentsScore["Suresh"]["Social"] = 90
	studentsScore["Ramesh"] = map[string]int{
		"Math":    80,
		"English": 70,
		"Social":  90,
		"Arts":    100,
	}

	for k, v := range studentsScore {
		fmt.Println(k)
		for k2, v2 := range v {
			fmt.Println(" ", k2, v2)
		}
	}

	studentsnestedScore := map[string]map[string]map[string]int{
		"Suresh": {
			"Midterm": {
				"Math":    80,
				"English": 70,
			},
			"Full term": {
				"Maths":   75,
				"English": 90,
			},
		},
		"Anand": {
			"Midterm": {
				"Math":    90,
				"English": 80,
			},
			"Full term": {
				"Maths":   75,
				"English": 90,
			},
		},
	}
	studentsnestedScore["Suresh"]["Midterm"]["Social"] = 90
	studentsnestedScore["Ramesh"] = map[string]map[string]int{
		"Midterm": {
			"Math":    80,
			"English": 70,
			"Social":  90,
			"Arts":    100,
		},
	}

	for k, v := range studentsnestedScore {
		fmt.Println(k)
		for k2, v2 := range v {
			fmt.Println(" ", k2)
			for k3, v3 := range v2 {
				fmt.Println("   ", k3, v3)
			}
		}
	}

	// -------------------------
	// 5. Nested Maps
	// -------------------------
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

	// -------------------------
	// 6. Struct as Map Key
	// -------------------------
	fmt.Println("6️⃣ Struct as Map Key")
	type Point struct{ X, Y int }
	points := make(map[Point]string)
	points[Point{1, 2}] = "Top Right"
	points[Point{-1, -2}] = "Bottom Left"
	PrintMap(points)
	fmt.Println()

	// -------------------------
	// 7. Sorted Key Iteration
	// -------------------------
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

