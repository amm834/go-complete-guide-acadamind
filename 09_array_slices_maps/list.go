package main

import "fmt"

func main() {
	var names [4]string
	var prices = []float64{10.2, 20.3, 20, 30}

	names = [4]string{
		"Mg Mg", "Hla Hla", "Aung Aung", "Moe Moe",
	}

	fmt.Println(prices)
	fmt.Println(names)

	fmt.Println(names[0])

	featuredPrices := prices[1:3]
	fmt.Println(featuredPrices)
}
