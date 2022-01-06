package main

import "fmt"

func main() {
	age := 18
	agePtr := &age

	fmt.Println(agePtr)

	*agePtr = 20
	println(age)

	doubledAge := double(agePtr)
	fmt.Printf("Doubled Age = %v\n", doubledAge)
	fmt.Printf("Changed Original Age = %v\n", age)

}

func double(age *int) int {
	*age = *age * 2
	return *age
}
