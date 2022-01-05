package main

import "fmt"

func main() {
	a := 10
	b := 20
	fmt.Println(sum(a, b))

	println("hello world")
}

func sum(a int, b int) int {
	return a + b
}

func println(x string) {
	fmt.Print(x)
}
