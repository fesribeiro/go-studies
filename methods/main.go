package main

import "fmt"

type user struct {
	name string
	age  int
}

func (u user) save() {
	fmt.Printf("Saving user: %s, Age: %d\n", u.name, u.age)
}

func (u user) isAdult() bool {
	return u.age >= 18
}

func (u *user) birthday() {
	u.age++
	fmt.Printf("Happy Birthday %s! You are now %d years old.\n", u.name, u.age)
}

func main() {
	u := user{name: "Alice", age: 17}
	u.save()
	if u.isAdult() {
		fmt.Println(u.name, "is an adult.")
	} else {
		fmt.Println(u.name, "is not an adult.")
	}

	fmt.Printf("Before birthday: %s is %d years old.\n", u.name, u.age)
	u.birthday()
}
