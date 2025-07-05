package main

import "fmt"

func main() {
	var var1 int = 10
	var var2 int = var1


	fmt.Println("Value of var1:", var1)
	fmt.Println("Value of var2:", var2)

	var1++ // Incrementing var1 does not change var2 because var2 is a copy of var1's value

	fmt.Println("After incrementing var1:", var1)
	fmt.Println("Value of var2 remains unchanged:", var2)

	var pointer1 *int = &var1

	var1++ // Incrementing var1 again will change the value pointed to by pointer1

	// this is dereferencing the pointer to get the value
	fmt.Println("Memory address:", pointer1, "Pointer to var1:", *pointer1)
	
}