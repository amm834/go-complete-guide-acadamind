package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

// Show message to user that can be helpful to understand what user have to do
func getUserInput(prompt string) string {
	fmt.Print(prompt)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Invalid User Input")
	}
	input = strings.ReplaceAll(input, "\n", "")
	return input
}

func Number(text string) int {
	number, err := strconv.Atoi(text)
	if err != nil {
		fmt.Println("Invalid String To Integer Conversion", err)
	}
	return number
}
