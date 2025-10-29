package main

import (
    "fmt"
    "math"
)

// Define Circle struct
type Circle struct {
    Radius float64
}
// Function to calculate the area of a Circle
func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

// Define Rectangle struct
type Rectangle struct {
    Length, Width float64
}
// Function to calculate the area of a Rectangle
func (r Rectangle) Area() float64 {
    return r.Length * r.Width
}


func main() {
    // Creating instances of Circle and Rectangle
    c := Circle{Radius: 5}           // Circle with radius 5
    r := Rectangle{Length: 4, Width: 6} // Rectangle 4x6

    // Directly calling the Area method for each type
    fmt.Printf("Area of Circle: %.2f\n", c.Area())      // Output: 78.54
    fmt.Printf("Area of Rectangle: %.2f\n", r.Area())   // Output: 24.00
}

