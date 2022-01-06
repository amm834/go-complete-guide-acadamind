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
	var newUser User
	fullName := getInput("Enter your name: ")
	birthDate := getInput("Enter yout birthday (DD/MM/YYYY): ")
	createdAt := time.Now()

	newUser = User{
		fullName, birthDate, createdAt,
	}
	fmt.Println(newUser)

}

func getInput(promptText string) string {
	fmt.Print(promptText)
	input, _ := reader.ReadString('\n')
	return input
}
