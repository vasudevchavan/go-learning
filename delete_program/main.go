package main

import (
	"fmt"
	"sort"
)

type colorstruct struct {
	name  string
	color string
}

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#000001ff",
	}

	kv := []colorstruct{}

	for k, v := range colors {
		kv = append(kv, colorstruct{name: k, color: v})
	}
	fmt.Println("Key-Value pairs in slice:", kv)

	sort.Slice(kv, func(i, j int) bool {
		return kv[i].name < kv[j].name
	})

	fmt.Println("Key-Value pairs in slice:", kv)
}
