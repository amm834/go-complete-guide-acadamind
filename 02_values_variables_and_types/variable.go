package main

import "fmt"

func main() {

	//var greetingText string
	//greetingText = "Hello World!"

	//var greetingText string = "Hello World!"

	greetingText := "Hello World!"

	luckyNumber := 19
	evenMoreLuckyNumber := luckyNumber + 1

	fmt.Println(greetingText)
	fmt.Println(luckyNumber)
	fmt.Println(evenMoreLuckyNumber)

	evenMoreLuckyNumber = luckyNumber * 2
	fmt.Println(evenMoreLuckyNumber)

	var newLuckyNumber float64 = float64(luckyNumber) / 3
	fmt.Println(newLuckyNumber)

	var defalutFloat float64 = 1.123456789123456789123456789
	var smallFloat float32 = 1.123456789123456789123456789
	fmt.Println(defalutFloat)
	fmt.Println(smallFloat)
}
