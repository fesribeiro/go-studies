package main

import "fmt"


func fibonacci(n uint) uint {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	position := uint(12) // Change this value to compute a different Fibonacci number

	for i:= uint(0); i <= position; i++ {
		fmt.Printf("Fibonacci of %d is %d\n", i, fibonacci(i))
	}
}
