package main

import (
	"fmt"
	"sort"

	"github.com/eiannone/keyboard"
)

// SortMenuByKeys sorts the coffee menu by keys and displays it
func SortMenuByKeys(coffee map[int]string) {
	var key []int
	for k := range coffee {
		key = append(key, k)
	}
	sort.Ints(key)
	fmt.Println("Welcome to the Coffee Shop!")
	fmt.Println("----------------------------")
	for _, k := range key {
		fmt.Printf("%d: %s\n", k, coffee[k])
	}
	fmt.Println("----------------------------")
	fmt.Println("Q or q to quit")
}

// CaptureUserInput captures user input and displays the selected coffee option

func CaptureUserInput(coffee map[int]string) {
	SortMenuByKeys(coffee)
	err := keyboard.Open()
	if err != nil {
		fmt.Println("Error opening keyboard:", err)
		return
	}
	defer keyboard.Close()

	char, _, err := keyboard.GetSingleKey()
	if err != nil {
		fmt.Println("Error reading key:", err)
		return
	}

	for char != 'q' && char != 'Q' {
		choice := int(char - '0')
		if coffeeName, exists := coffee[choice]; exists {
			fmt.Printf("You selected: %s\n", coffeeName)
		} else {
			fmt.Println("Invalid choice, please select a valid option.")
		}

		fmt.Println("\nPlease select again:")
		char, _, err = keyboard.GetSingleKey()
		if err != nil {
			fmt.Println("Error reading key:", err)
			return
		}
	}
}

// main function to run the program

func main() {

	coffee := make(map[int]string)
	coffee[1] = "Espresso"
	coffee[2] = "Latte"
	coffee[3] = "Cappuccino"
	coffee[4] = "Americano"
	coffee[5] = "Mocha"

	CaptureUserInput(coffee)
}
