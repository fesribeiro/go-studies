package main

import "fmt"

func func1() {
	// This function will be deferred
	println("Function 1 executed")
}

func func2() {
	// This function will also be deferred
	println("Function 2 executed")
}


func checkStudentIsApproved(n1,n2 float32) bool {
	defer fmt.Println("Grade check completed, the results will be printed below.")

	// This function checks if a student is approved based on their grades
	average := (n1 + n2) / 2
	if average >= 6 {
		fmt.Println("Student is approved with average:", average)
		return true
	} else {
		fmt.Println("Student is not approved with average:", average)
		return false
	}
}

func main() {
	// Defer the execution of func1 and func2
	defer func1()
	func2()




	fmt.Println("Checking student approval...", checkStudentIsApproved(7.5, 8.0))
	// Main function logic
	fmt.Println("Main function executed")

	// The deferred functions will be executed in reverse order
	// Output will be:
	// Main function executed
	// Function 2 executed
	// Function 1 executed
}
