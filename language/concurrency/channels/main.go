package main

import (
	"fmt"
	"time"
)

func write(text string, chanel chan string) {
	defer close(chanel) // Close the channel after sending all messages

	for i:= 0; i < 5; i++ {
		chanel <- text
		time.Sleep(1 * time.Second)
	}

}

func main() {
	channel := make(chan string)

	go write("Hello, World!", channel)

	fmt.Println("Waiting for messages...")
	
	// for {
	// 	message, open := <-channel // Wait for the first message
	// 	if !open {
	// 		break
	// 	}
	// 	fmt.Println(message)
	// }

	for message := range channel { // Wait for all messages
		fmt.Println(message)
	}
	fmt.Println("All messages received.")
	fmt.Println("This is a simple example of channels in Go.")
	fmt.Println("Channels allow goroutines to communicate with each other.")
}