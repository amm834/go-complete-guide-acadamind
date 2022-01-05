package io

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func GetInput(promptText string) float64 {
	fmt.Print(promptText)
	userInput, _ := reader.ReadString('\n')
	userInput = strings.ReplaceAll(userInput, "\n", "")
	convertedValue, _ := strconv.ParseFloat(userInput, 64)
	return convertedValue
}
