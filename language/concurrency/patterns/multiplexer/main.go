package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	channel := multiplexer(write("Hello"), write("World"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-channel)
	}
}

func multiplexer(channel1, channel2 <-chan string) <-chan string {
	ch := make(chan string)
	go func() {
		for {
			select {
			case msg1 := <-channel1:
				ch <- msg1
			case msg2 := <-channel2:
				ch <- msg2
			}
		}
	}()
	return ch
}

func write(text string) <-chan string {
	ch := make(chan string)
	go func() {
		for {
			ch <- fmt.Sprintf("Writing: %s", text)
			// Simulate some work
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()
	return ch
}