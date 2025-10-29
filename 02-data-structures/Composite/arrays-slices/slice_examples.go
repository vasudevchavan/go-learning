package main

import (
	"fmt"
	"golang.org/x/exp/slices"
)

func main() {

	// Creating slice:
	slice1 := make([]int, 0, 10)
	sl1 := []int{1, 2, 3}
	sl2 := make([]int, 5)
	sl3 := make([]int, 3, 10)

	// Appending elements
	fmt.Println("---------- Appending elements -----------")
	sl2 = append(sl2, 1, 2, 3, 4)
	fmt.Println("Slice")
	sliceLengthCapacity(sl2)
	slice1 = append(slice1, 1, 2, 3, 4, 5)
	fmt.Println("Slice")
	sliceLengthCapacity(slice1)
	slice1 = append(slice1, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 6)
	fmt.Println("Slice")
	sliceLengthCapacity(slice1)
	sl1 = append(sl1, sl3...)
	fmt.Println("Slice")
	sliceLengthCapacity(sl1)

	// Slicing a slice
	fmt.Println("---------- Slicing a slice -----------")
	fmt.Printf("From slice %v \nExtracting Range of slice %v \n", sl2, sl2[1:3])
	fmt.Printf("From slice %v \nExtracting Range of slice %v \n", sl2, sl2[:3])
	fmt.Printf("From slice %v \nExtracting Range of slice %v \n", sl2, sl2[:len(sl2)])

	// Example 2
	fmt.Println("---------- Slicing a slice -----------")
	slice11 := []int{2, 3, 4, 6, 8, 9, 10, 12, 14, 15, 20, 22}
	fmt.Printf("From slice %v \n", slice11)
	fmt.Printf("Extracting Range of slice %v \n", slice11[0:4])
	fmt.Printf("Extracting Range of slice %v \n", slice11[4:8])
	fmt.Printf("Extracting Range of slice %v \n", slice11[8:])

	// Copying Slice
	fmt.Println("---------- Copying Slice -----------")
	dstsl2 := make([]int, len(sl2))
	copy(dstsl2, sl2)
	fmt.Println("Copied Slice")
	sliceLengthCapacity(dstsl2)

	// Deleting elements from slice
	fmt.Println("---------- Deleting elements from slice -----------")
	fmt.Printf("Source Slice %v \n", dstsl2)
	// Removed all zeros
	dstsl2 = append(dstsl2[5:])
	fmt.Printf("Removed all zeros %v\n", dstsl2)
	// Removed all odd numbers
	dstsl2 = append(dstsl2[:2], dstsl2[3])
	dstsl2 = append(dstsl2[1:])
	fmt.Printf("Removed all odd numbers %v\n", dstsl2)

	// Iterating over slice
	fmt.Println("---------- Iterating over slice -----------")
	fmt.Println("Iterating over slice")
	for i := range sl2 {
		fmt.Printf("Value:%v\n", sl2[i])
	}
	for i, v := range sl2 {
		fmt.Printf("Index:%d - Value:%d\n", i, v)
	}

	//Multidimension Slice
	fmt.Println("---------- Multidimension Slice -----------")
	country1 := []string{"india", "new-delhi"}
	country2 := []string{"usa", "washington-dc"}
	country3 := []string{"russia", "moscow"}
	nations := [][]string{country1, country2, country3}

	fmt.Printf("Multidimension Slice\n%v\n", nations)

	// copying slice by referancing to slice directly
	fmt.Println("---------- copying slice by referancing to slice directly -----------")
	source := []int{1, 2, 3}
	destination := source
	// This change will impact destination slice since its referencing underlying array.
	fmt.Printf("Destination slice before copy %v\n", destination)
	source[0] = 100
	fmt.Printf("Source slice Modified %v\n", source)
	fmt.Printf("Destination slice also get affected %v\n", destination)
	if slices.Equal(source, destination) {
		fmt.Println("Source and destination slice are same (because they share the same underlying array)")
	}

	// coying slice using copy(destinatio,source)
	fmt.Println("---------- coying slice using copy(destinatio,source) -----------")
	source1 := []int{1, 2, 3}
	destination1 := make([]int, len(source1))
	copy(destination1, source1)
	source1[0] = 100
	if !slices.Equal(source1, destination1) {
		fmt.Println("Source and destination slice are different after updating source index")
	}

	// Reversing a slice
	fmt.Println("---------- Reversing a slice -----------")
	inputSlice := []int{1, 2, 3, 20, 5, 6, 7}
	fmt.Printf("slice to be reversed is %v\n", inputSlice)
	for i, j := 0, len(inputSlice)-1; i < j; i, j = i+1, j-1 {
		inputSlice[i], inputSlice[j] = inputSlice[j], inputSlice[i]
	}
	fmt.Printf("Reversed slice is %v", inputSlice)

	// Sorting a slice
	fmt.Println("---------- Sorting a slice -----------")
	sortSlice := []int{40, 2, 3, 20, 5, 6, 7}
	fmt.Println("Array used for Sorting \n",sortSlice)
	for i := 0; i <= len(sortSlice)-1; i++ {
		for j := i; j <= len(sortSlice)-1; j++ {
			if sortSlice[j] > sortSlice[i] {
				sortSlice[i], sortSlice[j] = sortSlice[j], sortSlice[i]
			}
		}
	}
	fmt.Println("Sorted array in desending\n",sortSlice)

}

func sliceLengthCapacity(m []int) {
	fmt.Println(m)
	fmt.Printf("Length:%d Capacity:%d \n", len(m), cap(m))
}
