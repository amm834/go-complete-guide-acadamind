package main

import (
	"fmt"
	"github.com/amm834/bmi/input"
	"github.com/amm834/bmi/maths"
)

func main() {
	weight := input.GetInput("Enter your weight(kg): ")
	height := input.GetInput("Enter your height(m): ")

	bmi := maths.CalculateBMI(weight, height)

	fmt.Printf("Your BMI: %.2f", bmi)
}
