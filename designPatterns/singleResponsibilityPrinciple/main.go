package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

var entryCount = 0

type Journal struct {
	entries []string
}

func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s",
		entryCount, text)
	j.entries = append(j.entries, entry)
	return entryCount
}

func (j *Journal) RemoveEntry(index int) {
	if index < 0 || index >= len(j.entries) {
		fmt.Println("Invalid index")
		return
	}
	fmt.Printf("Entry:%s is being removed\n", j.entries[index])
	slices.Delete(j.entries, index, index+1)
}

// Seperation of concerns
func (j Journal) String() string {
	return strings.Join(j.entries, "\n")
}

var LineSeparator = "\n"

func SaveToFile(j *Journal, filename string) {
	_ = os.WriteFile(filename, []byte(strings.Join(j.entries, LineSeparator)), 0644)
}

type Persistance struct {
	lineSeparator string
}

func (p *Persistance) SaveToFilePer(j *Journal, filename string) {
	_ = os.WriteFile(filename, []byte(strings.Join(j.entries, p.lineSeparator)), 0644)
}

func (j *Journal) LoadFromFile(filename string) {
	// ..
}

func main() {
	j := Journal{}
	j.AddEntry("India")
	j.AddEntry("Bharath")
	fmt.Printf("entries:\n%s\n", j.String())
	j.RemoveEntry(1)
	fmt.Printf("entries:\n%s\n", j.String())
	j.AddEntry("Bharath")
	SaveToFile(&j, "journal.txt")
	j.AddEntry("Hindustan")
	// p := Persistance{"\n"} two different ways to declare it
	p := Persistance{lineSeparator: "\n"}
	p.SaveToFilePer(&j, "journalPer.txt")
	fmt.Printf("entries:\n%s\n", j.String())
}
