package main

import (
	"fmt"
	"strings"
)

// OCP = Open for extension and closed for modification

type Color int

const (
	red Color = iota
	green
	blue
)

// Color is defined using iota and we are converting it to String
func (c Color) String() string {
	switch c {
	case red:
		return "Red"
	case green:
		return "Green"
	case blue:
		return "Blue"
	default:
		return "Unknown"
	}
}

// Size is defined using iota and we are converting it to String
func (c Size) String() string {
	switch c {
	case small:
		return "small"
	case medium:
		return "medium"
	case large:
		return "large"
	default:
		return "Unknown"
	}
}


type Size int

const (
	small Size = iota
	medium
	large
)

type Product struct {
	name  string
	color Color
	size  Size
}

type Filter struct {
}

func (f *Filter) FilterByColor(
	products []Product, color Color) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if v.color == color {
			result = append(result, &products[i])
		}
	}
	fmt.Printf("List of products with color %s:\n", color)
	return result
}

func (f *Filter) FilterBySize(
	products []Product, size Size) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if v.size == size {
			result = append(result, &products[i])
		}
	}
	fmt.Printf("List of products with size %s:\n", size)
	return result
}

func (f *Filter) FilterBySizeColor(
	products []Product, color Color, size Size) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if v.size == size && v.color == color {
			result = append(result, &products[i])
		}
	}
	fmt.Printf("List of products with color %s and size %s :\n", color, size)
	return result
}

func (p Product) String() string {
	return fmt.Sprintf("Product: %s, Color: %s  Size: %s", p.name, p.color.String(), p.size.String())
}

// interfaces
type Specification interface {
	IsSatisfied(p *Product) bool
}

type ColorSpecification struct {
	color Color
}

type SizeSpecification struct {
	size Size
}

type AndSpecification struct {
	first, second Specification
}

func (cs ColorSpecification) IsSatisfied(p *Product) bool {
	return p.color == cs.color
}

func (ss SizeSpecification) IsSatisfied(p *Product) bool {
	return p.size == ss.size
}

func (spec AndSpecification) IsSatisfied(p *Product) bool {
	return spec.first.IsSatisfied(p) && spec.second.IsSatisfied(p)
}

type BetterFilter struct{}

func (f *BetterFilter) Filter(
	products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if spec.IsSatisfied(&v) {
			result = append(result, &products[i])
		}
	}
	return result
}

func main() {
	apple := Product{"Apple", green, small}
	tree := Product{"Tree", green, large}
	water := Product{"BottleWater", blue, medium}
	products := []Product{apple, tree, water}
	productString := make([]string, len(products))

	// Listing all Products
	for i, p := range products {
		productString[i] = p.String()
	}
	result := strings.Join(productString, "\n")
	fmt.Println(result)

	colospec := ColorSpecification{color: blue}
	fmt.Println("We are validating color for products")
	for _, p := range products {
		if colospec.IsSatisfied(&p) {
			fmt.Printf("%s has valid color \n", p.name)
		}
	}

	// Filter Option which highlight Open Close Principle
	f := Filter{}
	for _, v := range f.FilterByColor(products, green) {
		fmt.Printf("- %s is green \n", v.name)
	}
	for _, v := range f.FilterBySize(products, small) {
		fmt.Printf("- %s is green \n", v.name)
	}
	for _, v := range f.FilterBySizeColor(products, green, small) {
		fmt.Printf("- %s is green \n", v.name)
	}

	fmt.Println("Green Product using Specifiction")
	greenSpec := ColorSpecification{green}
	bf := BetterFilter{}
	for _, v := range bf.Filter(products, greenSpec) {
		fmt.Printf(" - %s is green \n", v.name)
	}

	fmt.Println("Green Product large in size")
	greenAndLarge := AndSpecification{greenSpec, SizeSpecification{large}}
	for _, v := range bf.Filter(products, greenAndLarge) {
		fmt.Printf(" - %s is green \n", v.name)
	}

}
