package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

var reader = bufio.NewReader(os.Stdin)

type User struct {
	fullName  string
	birthDate string
	createdAt time.Time
}

func main() {
	var user User
	fullName := getInput("Enter your name: ")
	birthDate := getInput("Enter yout birthday (DD/MM/YYYY): ")
	createdAt := time.Now()

	user = User{
		fullName: fullName, birthDate: birthDate, createdAt: createdAt,
	}
}

func getInput(promptText string) string {
	fmt.Print(promptText)
	input, _ := reader.ReadString('\n')
	return input
}
