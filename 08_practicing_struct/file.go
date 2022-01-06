package main

import (
	"fmt"
	"os"
)

type Product struct {
	id, name, description, price string
}

func NewProduct(id, name, description, price string) *Product {
	return &Product{id, name, description, price}
}

func (product *Product) printDetail() {
	fmt.Printf("ID: %v\n", product.id)
	fmt.Printf("Name: %v\n", product.name)
	fmt.Printf("Description: %v\n", product.description)
	fmt.Printf("Price: %v", product.price)
	println("hello")
}

func (product *Product) store() {
	file, _ := os.Create(product.name)
	content := fmt.Sprintf("ID: %v\nName: %v\nDescription: %v\nPrice: %v$",
		product.id, product.name, product.description, product.price)
	file.WriteString(content)
	file.Close()
}

func main() {
	id := GetInput("Product ID: ")
	name := GetInput("Product Name: ")
	description := GetInput("Product Description: ")
	price := GetInput("Product Price: ")

	product := NewProduct(id, name, description, price)
	//product.printDetail()

	product.store()

}
