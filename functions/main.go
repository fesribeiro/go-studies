package main

import "fmt"

func sum(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func mathCalculation(n1, n2 int8) (int8, int8) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}

func main() {
	result := sum(10, 20)
	fmt.Println("The sum is:", result)

	_, sub := mathCalculation(30, 15)
	fmt.Println("The subtraction is:", sub)

	var f = func (text string) string {
		fmt.Println(text)
		return text
	}

	f("Hello, World!")
}