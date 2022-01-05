package main

import "fmt"

func main() {
	a := 10
	b := 20
	fmt.Println(sum(a, b))
}

func sum(a int, b int) int {
	return a + b
}
