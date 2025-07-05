package main

import "fmt"

type person struct {
	name string
	surname string
	age  int8
}

type student struct {
	person // Embedded struct
	course string
}

func main() {

	person1 := person{name: "John", surname: "Doe", age: 20}

	fmt.Println("Person 1:", person1)

	student1 := student{
		person1,
		"Computer Science",
	}

	fmt.Println("Student 1:", student1)
	fmt.Println("Student 1 Name:", student1.name)

}
