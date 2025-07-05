package main

import "fmt"

func main() {

	result := func (text string) string {
		return fmt.Sprintf("Anonymous function says: %s", text)
	}("This is an anonymous function") // Calling the anonymous function

	fmt.Println(result)

}
