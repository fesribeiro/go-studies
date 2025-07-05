package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}

func init() {
	fmt.Println("This is the init function, executed before main.")
}