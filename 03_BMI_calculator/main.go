package main

import (
	"fmt"
	"strconv"
	"strings"
)


func main() {
	fmt.Print("Enter your weight (kg): ")
	weightInput, _ := reader.ReadString('\n')

	fmt.Print("Enter your height (m): ")
	heightInput, _ := reader.ReadString('\n')

	weightInput = strings.ReplaceAll(weightInput, "\n", "")
	heightInput = strings.ReplaceAll(heightInput, "\n", "")

	weight,_ := strconv.ParseFloat(weightInput,64)
	height,_ := strconv.ParseFloat(heightInput,64)

	bmi := (weight / (height * height))
	fmt.Printf("Your BMI: %.2f",bmi)

}
