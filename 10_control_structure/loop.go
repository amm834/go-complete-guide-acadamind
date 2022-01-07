package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func main() {
	userChoice, err := getUserInput()
	if err != nil {
		fmt.Println("Invalid Choice Occured! Exiting!")
	}

	if userChoice == "1" {

	} else if userChoice == "2" {

	} else if userChoice == "3" {

	} else if userChoice == "4" {

	}

}

func getUserInput() (string, error) {
	fmt.Println("1) Sum up to X: ")
	fmt.Println("2) Sum factorial X: ")
	fmt.Println("3) Sum up manually entered numbers: ")
	fmt.Println("4) Sum up a list of entered numbers: ")

	var reader = bufio.NewReader(os.Stdin)
	userInput, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	if userInput == "1" ||
		userInput == "2" ||
		userInput == "3" ||
		userInput == "4" {

		return userInput, nil

	} else {
		return "", errors.New("INVALID INPUT")
	}

}
