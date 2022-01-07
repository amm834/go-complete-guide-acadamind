package main

import "fmt"

func main() {
	var names [4]string
	var prices = []float64{10, 20, 30, 40}

	names = [4]string{
		"Mg Mg", "Hla Hla", "Aung Aung", "Moe Moe",
	}

	fmt.Println(prices)
	fmt.Println(names)

	fmt.Println(names[0])

	featuredPrices := prices[2:]
	fmt.Println(featuredPrices)
}
