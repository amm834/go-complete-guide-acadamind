package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	input, err := getUserInput()

	if err != nil {
		fmt.Println(err)
		return
	}
	if input == "1" {
		calculateUntilNumber()
	} else if input == "2" {
		calculateFactorial()
	} else if input == "3" {
		calculateUserEnteredNumber()
	} else if input == "4" {
		calculateListOfNumbers()
	}
}

func showMenuList() {
	fmt.Printf(`
	Avaliable: 
		1) Sum up to x
		2) Sum factoial x
		3) Sum by Entering Numbers
		4) Sum by Entering List of Numbers
`)
}

func getUserInput() (string, error) {
	showMenuList()
	fmt.Print("Enter your choice: ")

	userInput, err := reader.ReadString('\n')

	userInput = strings.ReplaceAll(userInput, "\n", "")
	if err != nil {
		return "", err
	}

	if userInput == "1" || userInput == "2" || userInput == "3" || userInput == "4" {
		return userInput, nil
	} else {
		return "", errors.New("Invalid User Input Error")
	}
}

func calculateUntilNumber() {
	fmt.Print("Enter Your Number: ")
	inputNumber, err := getUserInputNumber()

	if err != nil {
		fmt.Println("Invalid Number Input")
		return
	}

	result := 0
	for i := 0; i <= inputNumber; i++ {
		result += i
	}

	fmt.Printf("Total Result: %d", result)

}

func calculateFactorial() {
	fmt.Print("Enter Your Number: ")
	inputNumber, err := getUserInputNumber()

	if err != nil {
		fmt.Println("Invalid Number Input")
		return
	}

	fmt.Printf("Total Number of factorial %d: %d", inputNumber, factorial(inputNumber))
}

func getUserInputNumber() (int, error) {
	numberInput, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	numberInput = strings.ReplaceAll(numberInput, "\n", "")

	number, err := strconv.Atoi(numberInput)
	if err != nil {
		return 0, err
	}
	return number, nil
}

func factorial(number int) int {
	if number <= 1 {
		return 1
	}
	return number * factorial(number-1)
}

func calculateUserEnteredNumber() {
	isEntering := true
	sum := 0
	for isEntering {
		fmt.Print("Enter your number: ")
		input, err := getUserInputNumber()
		isEntering = err == nil
		sum += input
	}
	fmt.Printf("Total Number: %d", sum)
}

func calculateListOfNumbers() {
	fmt.Println("Enter numbers list separate with comma (ex: 1,2,3)")
	fmt.Print("-> ")
	inputList, err := reader.ReadString('\n')
	inputList = strings.ReplaceAll(inputList, "\n", "")
	if err != nil {
		fmt.Println("Invalid Input")
		return
	}
	inputNumbers := strings.Split(inputList, ",")
	sum := 0
	// for each lol
	for _, value := range inputNumbers {
		number, err := strconv.Atoi(value)
		if err != nil {
			continue
		}
		sum += number
	}
	fmt.Printf("Total Numbers: %v", sum)

}
