package main

import "fmt"


func generic (interf interface{}) {
	fmt.Println("Generic function called with:", interf)
}

func main() {
	generic(42)                // Passing an int
	generic("Hello, World!")   // Passing a string
	generic(3.14)              // Passing a float64
	generic([]int{1, 2, 3})    // Passing a slice of ints
	generic(map[string]int{"a": 1, "b": 2}) // Passing a map
	fmt.Println("This is a simple example of generics in Go.")
	fmt.Println("Generics allow you to write functions that can accept any type.")
	fmt.Println("In this example, we have a generic function that prints the type and value of the argument.")
}