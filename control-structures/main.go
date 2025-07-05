package main

import "fmt"

func main() {
	number := 10

	if number > 0 {
		fmt.Println("The number is positive")
	} else if number < 0 {
		fmt.Println("The number is negative")
	} else {
		fmt.Println("The number is zero")
	}
	
	if otherNumber := number; otherNumber > 0 {
		fmt.Println("The other number is positive")
	} else {
		fmt.Println("The other number is not positive")
	}

}
