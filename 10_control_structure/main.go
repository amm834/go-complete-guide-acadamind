package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	ageInput := getInput("Enter your age: ")
	age, _ := strconv.Atoi(ageInput)
	if age >= 18 {
		fmt.Println("Welcome to club")
	}
}

func getInput(promptText string) string {
	fmt.Print(promptText)
	userInput, _ := reader.ReadString('\n')
	return strings.ReplaceAll(userInput, "\n", "")
}
