package main

import (
	"fmt"
	"time"
)

func write(text string) {
	for {
		fmt.Println(text)
		time.Sleep(1 * time.Second)
	}
}


func main() {
	go write("Hello, World!")
	write("Hello, Go!")
}