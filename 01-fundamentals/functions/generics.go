package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

type myNum interface {
	int | string
}

type myType interface {
	~int | ~string
}

type myConstraints   interface {
	constraints.Integer
}

func add2[T int | string](a, b T) (s T) {
	return a + b
}

func add2interface[T myNum](a, b T) (s T) {
	return a + b
}

func add2interfaceType[T myval](a, b T) (s T) {
	return a + b
}

func add2Constraints[T myConstraints](a, b T) (s T) {
	return a + b
}

type myval int

func main() {
	// Generics
	fmt.Println("addition result is %v", add2(1, 2))
	fmt.Println("addition result is %v", add2("1", "2"))

	// Generics using interface
	fmt.Println("addition result is %v", add2interface(1, 2))
	fmt.Println("addition result is %v", add2interface("1", "2"))

	// Type alias
	var aa, bb myval = -1, -2
	fmt.Println("addition result is %v", add2interfaceType(aa, bb))

		// Contraints alias
	var aaa, bbb myval = -1, -2
	fmt.Println("addition result is %v", add2interfaceType(aaa, bbb))

}
