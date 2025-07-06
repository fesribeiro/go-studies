package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type Dog struct {
	Name  string `json:"name"`
	Breed string `json:"breed"`
	Age   uint   `json:"age"`
}

func main() {
	dog := Dog{
		Name:  "Buddy",
		Breed: "Golden Retriever",
		Age:   3,
	}

	fmt.Println(dog)

	dogJSON, err := json.Marshal(dog)
	if err != nil {
		log.Fatal("Error marshalling dog:", err)
	}

	fmt.Println("Dog JSON:", dogJSON)
	fmt.Println("Dog JSON String:", bytes.NewBuffer(dogJSON))


	dog2 := map[string]string {
		"name":  "Buddy 2",
		"breed": "Golden Retriever",
	}

	dog2JSON, err := json.Marshal(dog2)
	if err != nil {
		log.Fatal("Error marshalling dog2:", err)
	}
	fmt.Println("Dog2 JSON:", dog2JSON)
	fmt.Println("Dog2 JSON String:", bytes.NewBuffer(dog2JSON))
}
