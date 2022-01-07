package main

import "fmt"

func main() {
	var names [4]string
	var prices = [4]float64{10, 20, 30, 40}

	names = [4]string{
		"Mg Mg", "Hla Hla", "Aung Aung", "Moe Moe",
	}

	fmt.Println(prices)
	fmt.Println(names)

	fmt.Println(names[0])

	featuredPrices := prices[2:]
	fmt.Println(featuredPrices)

	featuredName := names[:1]
	fmt.Println(len(featuredName), cap(featuredName))

	// Dynamic slices
	var products []string
	// Append and unpacking
	products = append(products, names[:]...)
	fmt.Println(products)
}
