package main

import (
    "fmt"
    "math"
)

// Define the Shape interface with Area method
type Shape interface {
    Area() (float64,string)
}

// Define Circle struct
type Circle struct {
    Radius float64
}
// Implement the Area method for Circle
func (c Circle) Area() (float64,string) {
    return math.Pi * c.Radius * c.Radius , "Circle"
}

// Define Rectangle struct
type Rectangle struct {
    Length, Width float64
}
// Implement the Area method for Rectangle
func (r Rectangle) Area() (float64,string) {
    return r.Length * r.Width, "Rectangle"
}

func main() {

    // First Method
    var s1 Shape = Circle{Radius: 5}       // Circle with radius 5
    var s2 Shape = Rectangle{Length: 4, Width: 6} // Rectangle 4x6
    aa1,ss1 := s1.Area()
    aa2,ss2 := s2.Area()
    fmt.Printf("Area of %s is : %.2f\n",ss1,aa1 )      // Output: 78.54
    fmt.Printf("Area of %s is: %.2f\n", ss2,aa2)   // Output: 24




    // Second Method using interface
    diffShape := []Shape{
        Circle{Radius: 5},
        Rectangle{Length: 4, Width: 6},
    }
    for _,v := range diffShape {
        area,shape := v.Area()
        fmt.Printf("Area of %s: %.2f\n",shape,area)
    }

}

