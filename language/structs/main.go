package main

import "fmt"


type address struct {
	street string
	city   string
}

type user struct {
	name  string
	age   uint8
	email string
	address  address // Embedded struct
}

func main() {

	address := address{
		street: "123 Main St",
		city:   "Springfield",
	}

	var user1 user
	user1.name = "John Doe"
	user1.age = 30
	user1.email = "john.doe@example.com"
	user1.address = address
	 
	fmt.Println("User 1:", user1)

	user2 := user{
		name:  "Jane Smith",
		age:   25,
		email: "jane.smith@example.com",
		address: address, // Using the same address struct
	}
	fmt.Println("User 2:", user2)

	user3 := user{name: "Alice Johnson", address: address}
	fmt.Println("User 3:", user3)
}
