package main

import "fmt"

func closure() func () { 
	text := "Inside closure function"

	functionClosure := func() {
		fmt.Println(text)
	}

	return functionClosure

}


func main() {
	text := "Inside main function"
	fmt.Println(text)

	newFuctionClosure := closure()
	newFuctionClosure() // Call the closure function to print the text from the closure
	fmt.Println("Closure function executed successfully")
}
