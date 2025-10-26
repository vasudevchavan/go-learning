package utils

import (
	"encoding/json"
	"fmt"
	"log"
)

// PrintMap prints key-value pairs of any map (using generics)
func PrintMap[K comparable, V any](m map[K]V) {
	fmt.Println("Printing map:")
	for k, v := range m {
		fmt.Printf("Key: %-10v | Value: %v\n", k, v)
	}
}

// PrintMapAsJSON prints any map as pretty JSON
func PrintMapAsJSON[K comparable, V any](m map[K]V) {
	data, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		log.Println("Error marshalling map to JSON:", err)
		return
	}
	fmt.Println("Map as JSON:\n", string(data))
}

// For simple string-to-string maps
func SimplePrintMap(m map[string]string) {
	fmt.Println("Map contents:", m)
}

func PrintMapKeyValue(m map[string]string) {
	for key, value := range m {
		fmt.Printf("Key: %-6s | Value: %s\n", key, value)
	}
}

func PrintAsJSON(m map[string]string) {
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		log.Println("Error marshalling map to JSON:", err)
	}
	fmt.Println("Map as JSON:", string(jsonData))
}

