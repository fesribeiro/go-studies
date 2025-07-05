package main

func calc(n1, n2 int) (sum int, diff int) {
	sum = n1 + n2
	diff = n1 - n2
	return
}

func main() {
	sum, diff := calc(10, 5)
	println("Sum:", sum)
	println("Difference:", diff)

	// Named return values can be used without explicitly returning them
	// This is useful for readability and clarity in the code
	sum, diff = calc(20, 15)
	println("Sum:", sum)
	println("Difference:", diff)
}
