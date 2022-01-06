package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

var reader = bufio.NewReader(os.Stdin)

type User struct {
	fullName  string
	birthDate string
	createdAt time.Time
}

func NewUser(fullName string, birthDate string) *User {
	createdAt := time.Now()
	user := User{fullName, birthDate, createdAt}
	return &user
}

func (user *User) printUser() {
	fmt.Printf("My name is %v (born on %v)", user.fullName, user.birthDate)
}

func main() {
	var newUser User
	fullName := getInput("Enter your name: ")
	birthDate := getInput("Enter yout birthday (DD/MM/YYYY): ")

	newUser = *NewUser(fullName, birthDate)
	newUser.printUser()
}

func getInput(promptText string) string {
	fmt.Print(promptText)
	input, _ := reader.ReadString('\n')
	cleaned := strings.ReplaceAll(input, "\n", "")
	return cleaned
}
