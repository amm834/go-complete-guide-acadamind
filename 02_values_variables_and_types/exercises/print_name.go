package main

import "fmt"

func main() {
	var firstName string = "Aung Myat"
	lastName := "Moe"

	var currentYear int = 2022
	birthYear := 2003
	currentAge := currentYear - birthYear

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(currentAge)

	currentYear = currentYear + 1
	nextAge := currentYear - birthYear
	fmt.Println(nextAge)
}
