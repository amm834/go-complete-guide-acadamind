package main

func main() {
	age := 18
	agePtr := &age

	println(agePtr)

	*agePtr = 19
	println(age)
}
