package main

import (
	"fmt"
	"reflect"
)

func main () {
	user := map[string]string{
		"name": "Felipe",
		"age":  "30",
		"city": "São Paulo",
	}

	user2 := map[string]map[string]string{
		"Felipe": {
			"age":  "30",
			"city": "São Paulo",
		},
	}

	fmt.Println("User Map:", user, reflect.TypeOf(user))
	fmt.Println("User2 Map:", user2, reflect.TypeOf(user2))

	delete(user, "age")
	fmt.Println("User Map after deletion:", user)

	user["age"] = "31"
	fmt.Println("User Map after update:", user)
}
