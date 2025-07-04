package main

import "fmt"


func main() {
	var var1 int = 10
	var2 := 20


	fmt.Println(var1, var2)
	
	var (
		var3 int = 30
		var4 int = 40
	)

	fmt.Println(var3, var4)

	var5, var6 := 50, 60
	fmt.Println(var5, var6)

	const const1 string = "Constant Value 1"
	fmt.Println(const1)
}