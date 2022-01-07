package main

import "fmt"

func main() {
	var names [4]string
	var prices = []float64{10.2, 20.3, 20, 20}

	names = [4]string{
		"Mg Mg", "Hla Hla", "Aung Aung", "Moe Moe",
	}

	prices[2] = 400
	fmt.Println(prices)
	fmt.Println(names)
}
