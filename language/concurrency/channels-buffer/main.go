package main

import (
	"fmt"
)

func main() {
	channel := make(chan string, 2)

	channel <- "Hello, World!" // Send a message to the channel

	message := <-channel // Receive the message from the channel
	fmt.Println(message)

}