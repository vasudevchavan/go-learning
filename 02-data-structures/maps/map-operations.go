package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func printMapKeyValue(m map[string]string) {
	for key, value := range m {
		fmt.Printf("Key: %-6s, | Value: %s\n", key, value)
	}
}

func simplePrintMap(m map[string]string) {
	fmt.Println("Map contents:", m)
}

func printAsJson(m map[string]string) {
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		log.Println("Error marshalling map to JSON:", err)
	}
	fmt.Println("Map as JSON:", string(jsonData))
}

func main() {

	// Way of declaring a Maps
	var empty_colors map[string]string
	empty_colors1 := make(map[string]string)
	empty_colors2 := make(map[string]string, 10)

	fmt.Println("Empty map:", empty_colors, empty_colors1, empty_colors2)

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#000001ff",
	}

	// Adding a new key-value pair
	colors["white"] = "#ffffff"
	colors["black"] = "#000000"

	printMapKeyValue(colors)
	simplePrintMap(colors)
	printAsJson(colors)

}
