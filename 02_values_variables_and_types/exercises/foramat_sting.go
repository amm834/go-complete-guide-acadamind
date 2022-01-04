package main

import (
	"fmt"
	"math"
)

func main() {
	const PI = math.Pi
	radius := 5
	circumference := 2 * PI * float64(radius)

	fmt.Printf("For a radius of %v, the circle circumference is %.2f", radius, circumference)
}
