package main

import "fmt"


func sum (numbers ...int) int {
	fmt.Println("Numbers:", numbers)
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func write(text string, numbers ...int) {
	fmt.Println(text)
	for _, number := range numbers {
		fmt.Println("Number:", number)
	}
}

func main () {
	fmt.Println("Sum:", sum(1, 2, 3, 4, 5))
	write("Numbers:", 10, 20, 30, 40, 50)
}
