package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func GetInput(promptText string) string {
	fmt.Print(promptText)
	input, _ := reader.ReadString('\n')
	cleanedInput := strings.ReplaceAll(input, "\n", "")
	return cleanedInput
}
