package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
}

type ContactInfo struct {
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

type Employee struct {
	Person
	ContactInfo
	EmployeeID string `json:"employeeId"`
	Department string `json:"department"`
}

func (e Employee) convertToJSON() (string, error) {
	data, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		log.Println("Error marshalling employee data:", err)
		return "", err
	}
	return string(data), nil
}

func (e *Employee) UpdateFirstName(newFirstName string) {
	e.FirstName = newFirstName
}

func (e *Employee) UpdateLastName(newLastName string) {
	(*e).LastName = newLastName
}

func (e *Employee) NewUpdateFirstName(newFirstName string) {
	e.FirstName = newFirstName
}

func (e *Employee) NewUpdateLastName(newLastName string) {
	(*e).LastName = newLastName
}

func main() {

	printFlag := false
	// Direct initialization of a struct
	person := []Person{}

	person = append(person, Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
	})

	// Initializing a struct using a pointer
	ptrperson1 := &Person{
		FirstName: "Alice",
		LastName:  "Johnson",
		Age:       28,
	}

	ptrperson2 := new(Person)
	ptrperson2.FirstName = "Bob"
	ptrperson2.LastName = "Brown"
	ptrperson2.Age = 35

	// Initializing an array of structs
	arrayPerson := [2]Person{
		{"Charlie", "Davis", 22},
		{"Eve", "White", 29},
	}

	// Initializing a slice of structs
	slicePerson := []Person{
		{"Frank", "Black", 33},
		{"Grace", "Green", 27},
	}

	ptrSlicePerson := []*Person{
		&Person{"Hank", "Blue", 40},
		&Person{"Ivy", "Purple", 31},
	}

	// *** How to Print the structs
	// fmt.Println("Printing Person Details:", person)
	// fmt.Printf("Printing Pointer Person 1:%+v", *ptrperson1)
	// fmt.Printf("\nPrinting Pointer Person 2:%+v", *ptrperson2)
	// fmt.Println("\nPrinting Array of Persons:")
	// for _, p := range arrayPerson {
	// 	fmt.Printf("%+v\n", p)
	// }
	// fmt.Printf("Printing slice of Persons\n")
	// for _, p := range slicePerson {
	// 	fmt.Printf("%+v\n", p)
	// }
	// fmt.Printf("Printing pointer slice of Persons\n")
	// for _, p := range ptrSlicePerson {
	// 	fmt.Printf("%+v\n", *p)
	// }

	allPersons := []Person{}
	allPersons = append(allPersons, person...)
	allPersons = append(allPersons, *ptrperson1, *ptrperson2)
	allPersons = append(allPersons, arrayPerson[:]...)
	allPersons = append(allPersons, slicePerson...)
	for _, p := range ptrSlicePerson {
		allPersons = append(allPersons, *p)
	}

	// fmt.Println("\nAll Persons Combined:")
	// for _, p := range allPersons {
	// 	fmt.Printf("%v\n", p)
	// }

	data, err := json.MarshalIndent(allPersons, "", "  ")
	// Explaination starts here
	//JSON string.
	// The second and third arguments ("", "  ") define:
	// ⁠"": no prefix before each line
	// ⁠"  ": two spaces for indentation
	// data is a []byte containing the JSON
	// err captures any error during conversion
	// Explaination ends here

	if err != nil {
		fmt.Println("Error marshalling data:", err)
	}
	if printFlag {
		fmt.Println(string(data))
	}

	// Embeded structs example
	employee := Employee{
		Person: Person{
			FirstName: "David",
			LastName:  "Smith",
			Age:       45,
		},
		ContactInfo: ContactInfo{
			Email:   "davitsmith@gmail.com",
			Phone:   "123-456-7890",
			Address: "123 Elm St, Springfield",
		},
		EmployeeID: "E12345",
		Department: "Engineering",
	}
	fmt.Println("\nEmployee Details:", employee)

	// Print the employee details in JSON format

	// employeeData, err := json.MarshalIndent(employee, "", "  ")
	// if err != nil {
	// 	fmt.Println("Error marshalling employee data:", err)
	// }
	// fmt.Println(string(employeeData))
	employee.UpdateFirstName("Rahul")
	employee.UpdateLastName("Dravid")
	employeeJSON, err := employee.convertToJSON()
	fmt.Println("\nEmployee JSON:", employeeJSON)

	newEmployee := &employee
	newEmployee.UpdateFirstName("Sachin")
	newEmployee.UpdateLastName("Tendulkar")

	employeeJSON, err = employee.convertToJSON()
	fmt.Println("\nEmployee JSON:", employeeJSON)

	employee.NewUpdateFirstName("Virat")
	employee.NewUpdateLastName("Kohli")

	employeeJSON, err = employee.convertToJSON()
	fmt.Println("\nEmployee JSON:", employeeJSON)

}
