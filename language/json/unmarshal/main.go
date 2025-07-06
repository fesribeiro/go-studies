package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Dog struct {
	Name  string `json:"name"`
	Breed string `json:"breed"`
	Age   uint   `json:"-"` // Omit the Age field from JSON
}

func main() {
	dogJSON := `{"breed":"Golden Retriever","name":"Buddy"}`

	var dog Dog

	if erro := json.Unmarshal([]byte(dogJSON), &dog); erro != nil {
		log.Fatal("Error unmarshalling dog:", erro)
	}

	dog2JSON := `{"name":"Buddy 2","breed":"Golden Retriever"}`

	dog2 := make(map[string]string)

	if erro := json.Unmarshal([]byte(dog2JSON), &dog2); erro != nil {
		log.Fatal("Error unmarshalling dog2:", erro)
	}

	fmt.Println(dog)
	fmt.Println(dog2)
}
