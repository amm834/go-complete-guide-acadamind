package main

import (
	"fmt"
	"math/rand"
)

func main() {
	a, b := printNumbers()
	sum := sum(a, b)
	fmt.Println(sum)
	println("hello world")
}

func sum(a int, b int) (sum int) {
	sum = a + b
	return
}

func println(text string) {
	fmt.Print(text)
}

func printNumbers() (int, int) {
	rand1 := rand.Intn(10)
	rand2 := rand.Intn(10)
	return rand1, rand2
}
