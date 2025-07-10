package main

import (
	"fmt"
	"strings"
)

// Aggregators struct

type myStruct struct {
	Luxury        bool
	Make          string
	Model         string
	Year          int
	NumberOfTyres int
	BucketSeat    bool
}

func UpdateCar(c *myStruct, make string, model string, year int) {
	c.Make = make
	c.Model = model
	c.Year = year
}

func main() {
	a := true
	b := false
	b = !a
	fmt.Printf("head %v or tail %v is a == b %v", a, b, a == b)

	i := 10
	var j int = 11
	var f float32 = 11.5
	fmt.Printf("adding i %v + % v and result is %v", i, j, i+j)
	fmt.Printf("float value", f)

	var tell string = "I am in india"
	fmt.Printf("do you live in india %v", strings.HasSuffix(tell, "india"))

	var aa int = 100

	switch aa % 3 {
	case 0:
		fmt.Println("divisible by 3")
	default:
		fmt.Println("not divisible")
	}

	// Looping using index

	for _, i := range []int{10, 11, 12, 13, 14, 15, 16, 17, 18} {
		fmt.Println(i)
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	ab := 2
	for ab < 10 {
		fmt.Println(ab)
		ab = ab + 2
	}

	// Aggregators array

	var myArrayInt [5]int
	myArrayInt[0] = 10
	myArrayInt[1] = 20
	myArrayInt[2] = 30
	myArrayInt[3] = 40
	myArrayInt[4] = 50
	fmt.Println("my array is ", myArrayInt)

	var myArrayString [5]string
	myArrayString[0] = "one"
	myArrayString[1] = "two"
	myArrayString[2] = "three"
	myArrayString[3] = "four"
	myArrayString[4] = "five"
	fmt.Println("my string array is ", myArrayString)

	var myCar myStruct
	myCar.Luxury = true
	myCar.Make = "BMW"
	myCar.Model = "X5"
	myCar.Year = 2023
	myCar.NumberOfTyres = 4
	myCar.BucketSeat = true
	fmt.Println("my car is ", myCar)

	var myCars []myStruct
	myCars = append(myCars, myCar)
	myCars = append(myCars, myStruct{Luxury: false, Make: "Toyota", Model: "Corolla", Year: 2020, NumberOfTyres: 4, BucketSeat: false})
	fmt.Println("my cars are ", myCars)

	// Pointers

	myNewCar := &myStruct{
		Luxury:        true,
		Make:          "Tesla",
		Model:         "Model S",
		Year:          2023,
		NumberOfTyres: 4,
		BucketSeat:    true}

	fmt.Println("my new car is ", myNewCar)
	UpdateCar(myNewCar, "Tesla", "Model X", 2024)
	fmt.Println("my new car is ", myNewCar)

}
