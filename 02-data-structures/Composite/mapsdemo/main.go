package main

import (
	"fmt"

	"mapsdemo/utils"
)

func main() {
	runSimpleNested()
	runDeepNested()
	runDedup()
	runCharFrequency()
	runWordFrequency()

	utils.BasicMapOperations()
	utils.ArrayToMapConversion()
	utils.DuplicatesHandling()
	utils.MapsOfSlices()
	utils.NestedMaps()
	utils.StructAsMapKey()
	utils.SortedKeyIteration()

	// Existing examples
	utilsExampleRunner()

	// 🚀 New Advanced Algorithms
	utils.TestAdvancedMapOperations()
}

func utilsExampleRunner() {
	runSimpleNested()
	runDeepNested()
	runDedup()
	runCharFrequency()
	runWordFrequency()
}

func runSimpleNested() {
	fmt.Println("=== Simple Nested Maps ===")
	s := utils.NewStudentScores()
	s.AddScore("Suresh", "Math", 80)
	s.AddScore("Suresh", "English", 70)
	s.AddScore("Suresh", "Social", 90)
	s.AddScore("Anand", "Math", 90)
	s.AddScore("Ramesh", "Arts", 100)

	jsonStr, _ := s.ToJSON()
	fmt.Println(jsonStr)
}

func runDeepNested() {
	fmt.Println("\n=== Deep Nested Maps ===")
	d := utils.NewDeepScores()
	d.AddScore("Suresh", "Midterm", "Math", 80)
	d.AddScore("Suresh", "Midterm", "English", 70)
	d.AddScore("Suresh", "Full term", "Maths", 75)
	d.AddScore("Ramesh", "Midterm", "Arts", 100)

	jsonStr, _ := d.ToJSON()
	fmt.Println(jsonStr)
}

func runDedup() {
	fmt.Println("\n=== Remove Duplicate from Array ===")
	arr := []string{"a", "b", "c", "d", "c", "a", "d", "c", "g", "h"}
	unique, counts := utils.RemoveDuplicateArray(arr)
	fmt.Println("Unique values:", unique)
	fmt.Println("Counts:", counts)

	samplemap := map[string]int{
		"Go":   1,
		"C":    2,
		"C++":  3,
		"Java": 3,
	}
	samplemap["zig"] = 4

	fmt.Println("\n=== Reverse Maps ignoring Duplicate ===")
	fmt.Println(utils.ReverseMapIgnoringDup(samplemap))
	fmt.Println("\n=== Reverse Maps Considering Duplicate ===")
	fmt.Println(utils.ReverseMapConsiderDup(samplemap))
}

func runCharFrequency() {
	fmt.Println("\n=== Character Frequency ===")
	str := "asfaesaereaeddsafsawwqqwewr"
	counts := utils.RemoveDuplicateChars(str)
	for ch, count := range counts {
		if count > 1 {
			fmt.Printf("%c occurred %d times\n", ch, count)
		}
	}
}

func runWordFrequency() {
	fmt.Println("\n=== Word Frequency ===")
	words := []string{"go", "is", "fun", "go", "is", "fast"}
	counts := utils.FrequencyCounter(words)
	for word, count := range counts {
		fmt.Printf("%s occurred %d times\n", word, count)
	}
}
