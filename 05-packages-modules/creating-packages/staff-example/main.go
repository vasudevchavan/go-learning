package main

import (
	"fmt"
	"myapp/staff"
)

var addEmployees = []staff.Employee{
	{
		FirstName: "John",
		LastName:  "Doe",
		FullTime:  true,
		Salary:    50000,
	},
	{
		FirstName: "Jane",
		LastName:  "Smith",
		FullTime:  false,
		Salary:    40000,
	},
	{
		FirstName: "Alice",
		LastName:  "Johnson",
		FullTime:  true,
		Salary:    60000,
	},
	{FirstName: "Bob",
		LastName: "Brown",
		FullTime: false,
		Salary:   45000,
	},
	{
		FirstName: "Charlie",
		LastName:  "Davis",
		FullTime:  true,
		Salary:    70000,
	},
}

var addEmployees1 = staff.Employee{
	FirstName: "Eve",
	LastName:  "White",
	FullTime:  false,
	Salary:    48000,
}

func main() {
	myStaff := staff.Office{
		AllStaff: addEmployees,
	}
	myStaff.AllStaff = append(myStaff.AllStaff, addEmployees1)

	fmt.Println(myStaff.PrintAllStaff1())

	// Print OverPaid Employees
	fmt.Println("Overpaid Employees:", myStaff.OverPaidEmployees())
	fmt.Println("UnderPaid Employees:", myStaff.UnderPaidEmployees())

}
