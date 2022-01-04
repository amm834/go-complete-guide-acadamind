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
}
