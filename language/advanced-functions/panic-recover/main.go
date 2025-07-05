package main

import "fmt"


func recoverFromPanic() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from panic: successfully", r)
	}
}

func checkStudentIsApproved(n1, n2 float32) bool {
	defer recoverFromPanic()
	grade := (n1 + n2) / 2

	if grade > 6 {
		return true
	} else if grade < 6 {
		return false
	} else {
		panic("Grade cannot be exactly 6, please check the input values.")
	}
}


func main() { 
	fmt.Println("Checking student approval...")
	fmt.Println("Student is approved:", checkStudentIsApproved(8,6))
}