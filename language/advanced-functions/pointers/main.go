package main


func changeSign(number int) int {
	return number * -1
}

func changeSignPointer(number *int) {
	*number = *number * -1
}

func main() { 
	number := 10
	// Call the changeSign function to change the sign of the number
	newNumber := changeSign(number)
	println("Changed sign number:", newNumber)

	changeSignPointer(&number)
	println("Changed sign number with pointer:", number)

	println("Original number:", number)


}