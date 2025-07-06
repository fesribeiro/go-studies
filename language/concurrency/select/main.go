package main

import (
	"fmt"
	"time"
)

func main() {
	channel1, channel2 := make(chan string), make(chan string) // Create a new channel of type string

	go func() {
		for {
			time.Sleep(time.Millisecond * 500) // Simulate some work
			channel1 <- "channel1" // Send a message to channel1
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2) // Simulate some work
			channel2 <- "channel2" // Send a message to channel2
		}
	}()

	for {
		select {
			case messageFromChannel1 := <-channel1: // If a message is received from channel
				fmt.Println("Received from", messageFromChannel1) // Print the message from channel1
			case messageFromChannel2 := <-channel2: // If a message is received from channel2
				fmt.Println("Received from", messageFromChannel2) // Print the message from channel2
		}
	}

}