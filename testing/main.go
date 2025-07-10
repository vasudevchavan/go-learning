package main

import (
	"fmt"
)

// Product struct holds information about a Product
type Product struct {
	name  string
	color string
	size  string
}

func FilterWrongProduct(products []*Product, condition func(*Product) bool) []*Product {
	var result []*Product
	for _, pd := range products {
		if condition(pd) {
			fmt.Printf("Product Name:%s of color:%s and size:%s \n", pd.name, pd.color, pd.size)
			result = append(result, pd)
		}
	}
	return result
}

func IsWrongProduct(p *Product) bool {
	return p.color == "" && p.size == ""
}

func main() {
	apple := &Product{"Apple", "red", "small"}
	phone := &Product{"Phone", "blue", "large"}
	dummy := &Product{"Dummy", "", ""}

	products := []*Product{apple, phone, dummy}

	wrongproduct := FilterWrongProduct(products, IsWrongProduct)

	for i, p := range wrongproduct {
		fmt.Printf("%d Product Name:%s of color:%s and size:%s \n", i, p.name, p.color, p.size)
	}

}
