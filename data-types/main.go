package main

import (
	"errors"
	"fmt"
)


func main() {
	number := 1000000000000000000

	fmt.Println("The value of number is:", number)

	var number2 uint32 = 10000
	fmt.Println("The value of number2 is:", number2)


	//alias
	// int32 is an alias for int32
	var number3 rune = 123456
	fmt.Println("The value of number3 is:", number3)

	// byte is an alias for uint8
	var number4 byte = 255
	fmt.Println("The value of number4 is:", number4)


	var number5 float32 = 3.14
	fmt.Println("The value of number5 is:", number5)

	var number6 float64 = 3.141592653589793
	fmt.Println("The value of number6 is:", number6)

	var str string = "Hello, World!"
	fmt.Println("The value of str is:", str)

	str2 := "Hello, Go!"
	fmt.Println("The value of str2 is:", str2)

	char := 'A'
	fmt.Println("The value of char is:", char)

	var text int
	fmt.Println("The value of text is:", text)


	var booleano bool = true
	fmt.Println("The value of booleano is:", booleano)

	var err error = errors.New("This is an error")
	fmt.Println("The value of error is:", err)
}