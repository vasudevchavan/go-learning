package staff

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	FirstName string
	LastName  string
	FullTime  bool
	Salary    int
}

type Office struct {
	AllStaff []Employee
}

func (o *Office) PrintAllStaff() {
	dataJson, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling data:", err)
		return
	}
	fmt.Println("JSON Data:", string(dataJson))
}

func (o *Office) PrintAllStaff1() []Employee {
	return o.AllStaff
}

func (o *Office) OverPaidEmployees() []Employee {
	var overPaid []Employee
	for _, emp := range o.AllStaff {
		if emp.FullTime && emp.Salary > 60000 {
			overPaid = append(overPaid, emp)
		}
	}
	return overPaid
}

func (o *Office) UnderPaidEmployees() []Employee {
	var underPaid []Employee
	for _, emp := range o.AllStaff {
		if !emp.FullTime && emp.Salary < 50000 {
			underPaid = append(underPaid, emp)
		}
	}
	return underPaid
}
