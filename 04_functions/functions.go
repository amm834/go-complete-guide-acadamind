package main

import (
	"fmt"
	"math/rand"
)

func main() {
	a, b := printNumbers()
	fmt.Println(sum(a, b))
	println("hello world")
}

func sum(a int, b int) int {
	return a + b
}

func println(text string) {
	fmt.Print(text)
}

func printNumbers() (int, int) {
	rand1 := rand.Intn(10)
	rand2 := rand.Intn(10)
	return rand1, rand2
}
