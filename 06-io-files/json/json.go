package main

import (
	"fmt"
	// "strconv"
	"encoding/json"
	"log"
)

func main() {
	xs := []string{"in", "my", "younger", "and", "more", "vulnerable", "years", "my", "father"}
	newMap, countMap := make(map[string]int),make(map[string]int)
	for _, v := range xs {
		countMap[v]++
	}	

	// Convert Map to JSON
	copyjson := PrintAsJSON(countMap)
	fmt.Println(string(copyjson))

  // Convert Json to Map
	err := json.Unmarshal(copyjson,&newMap)
	if err != nil {
		log.Println("Error marshalling map to JSON:", err)
	}

	for k,v:= range newMap {
		fmt.Printf("key:%s - value:%d \n",k,v)
	}
}

func PrintAsJSON(m map[string]int) ([]byte){
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		log.Println("Error marshalling map to JSON:", err)
	}
	fmt.Println("Map as JSON:", string(jsonData))
	return jsonData
}
