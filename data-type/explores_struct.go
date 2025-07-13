package main

import "fmt"

type Person struct {
	Name  string
	Skill []string
}

func main() {

	// Declaration of a struct
	var person1 Person
	person1.Name = "John"
	person1.Skill = []string{"Go", "Python", "Java"}
	fmt.Println("Person 1:", person1)

	// Declaration of a slice of structs to add multiple persons
	var person2 []Person
	person2 = append(person2, person1)
	person2 = append(person2, Person{Name: "Alice", Skill: []string{"JavaScript", "C++"}})
	person2 = append(person2, Person{Name: "Bob", Skill: []string{"Ruby", "PHP"}})
	fmt.Println("Person 2:", person2)
}
