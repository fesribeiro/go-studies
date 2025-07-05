package main

import (
	"fmt"
	"time"
)

func main() {
	// i := 0

	// for i < 5 {
	// 	i++
	// 	fmt.Println("Value of i:", i)
	// 	time.Sleep(time.Second)
	// }

	// for j := 0; j < 5; j++ {
	// 	fmt.Println("Value of j:", j)
	// 	time.Sleep(time.Second)
	// }

	names := [3]string{"Alice", "Bob", "Charlie"}

	for index, name := range names {
		fmt.Printf("Index: %d, Name: %s\n", index, name)
		time.Sleep(time.Second)
	}

	for index, letter := range "Hello" {
		fmt.Printf("Index: %d, Letter: %s\n", index, string(letter))
		time.Sleep(time.Second)
	}

	user := map[string]string{
		"name":    "Alice",
		"country": "Wonderland",
		"age":     "25",
	} 

	for key, value := range user {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
		time.Sleep(time.Second)
	}

	for {
		fmt.Println("This will run forever. Press Ctrl+C to stop.")
		
	}

}