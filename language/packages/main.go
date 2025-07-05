package main

import (
	"fmt"
	"module/helpers"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(helpers.Greet("Go Developer"))
	err := validateEmail("test@example.com")
	if err != nil {
	fmt.Println("Invalid email format:", err)
	} else {
		fmt.Println("Valid email format.")
	}
	fmt.Println(helpers.Farewell("Go Developer"))
	fmt.Println("This is a simple Go application using packages.")
}

func validateEmail(email string) error {
	return checkmail.ValidateFormat(email)
}