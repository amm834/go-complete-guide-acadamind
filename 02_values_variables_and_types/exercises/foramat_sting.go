package main

import (
	"fmt"
	"math"
)

func main() {
	PI := math.Pi
	radius := 5
	circumference := 2 * PI * float64(radius)

	fmt.Printf("For a radius of %v, the circle circumference is %.2f", radius, circumference)
}
