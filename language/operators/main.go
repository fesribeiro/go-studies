package main

import "fmt"


func main() {
	// operations
	// +
	// -
	// *
	// /
	// %
	// ==
	// !=
	// <
	// >
	// <=
	// >=
	// &&
	// ||
	// !


	sum := 5 + 3
	diff := 10 - 4
	product := 6 * 7
	quotient := 20 / 4
	remainder := 10 % 3
	equal := 5 == 5
	notEqual := 5 != 3
	lessThan := 3 < 5
	greaterThan := 7 > 2
	lessThanOrEqual := 4 <= 4
	greaterThanOrEqual := 8 >= 6
	andOperation := (5 > 3) && (2 < 4)
	orOperation := (5 < 3) || (2 < 4)
	notOperation := !(5 > 3)
	fmt.Println("Sum:", sum)
	fmt.Println("Difference:", diff)
	fmt.Println("Product:", product)
	fmt.Println("Quotient:", quotient)
	fmt.Println("Remainder:", remainder)
	fmt.Println("Equal:", equal)
	fmt.Println("Not Equal:", notEqual)
	fmt.Println("Less Than:", lessThan)
	fmt.Println("Greater Than:", greaterThan)
	fmt.Println("Less Than or Equal:", lessThanOrEqual)
	fmt.Println("Greater Than or Equal:", greaterThanOrEqual)
	fmt.Println("And Operation:", andOperation)
	fmt.Println("Or Operation:", orOperation)
	fmt.Println("Not Operation:", notOperation)
	fmt.Println("All operations completed successfully.")

	number := 42
	number++
	number += 5

	fmt.Println("Incremented Number:", number)

	number--
	number -= 3
	fmt.Println("Decremented Number:", number)

	fmt.Println("End of operations.")
	fmt.Println("This is a simple demonstration of basic operations in Go.")
}
