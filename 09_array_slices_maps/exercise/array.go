package main

import "fmt"

type Product struct {
	name  string
	price float64
}

func NewProduct(name string, price float64) Product {
	return Product{
		name, price,
	}
}

func main() {
	var hobbies = [3]string{
		"Coding", "Eating", "Watching",
	}
	fmt.Println(hobbies)

	// standalone
	firstElement := hobbies[0]
	fmt.Println(firstElement)

	// second and third
	modifiedHobbies := hobbies[1:]
	fmt.Println(modifiedHobbies)

	newHobbies := hobbies[:2]
	fmt.Println(newHobbies)

	combinedHobbies := append(newHobbies, hobbies[:1]...)
	fmt.Println("Combined ", combinedHobbies)

	productOne := NewProduct("Cocacola", 20.22)
	productTwo := NewProduct("Waifer", 30.2)

	var products = []Product{productOne, productTwo}

	productThree := NewProduct("Snack", 20.2)
	products = append(products, productThree)

	fmt.Println(products, len(products))

}
