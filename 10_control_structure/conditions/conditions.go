package conditions

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
	age, err := strconv.Atoi(ageInput)

	if err != nil {
		fmt.Println("Invalid Input")
		return
	}

	if (age >= 30 && age <= 50) || age >= 60 {
		fmt.Println("You are eligible for senior job")
	} else if age >= 50 {
		fmt.Println("The best age")
	} else if age >= 18 {
		fmt.Println("Welcome to club")
	} else {
		fmt.Println("Sorry, we don't accept a child")
	}
}

func getInput(promptText string) string {
	fmt.Print(promptText)
	userInput, _ := reader.ReadString('\n')
	return strings.ReplaceAll(userInput, "\n", "")
}
