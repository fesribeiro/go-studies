package main

import (
	"fmt"
	"time"
)


func write(text string) <- chan string {
	ch := make(chan string)
	go func() {
		for {
			ch <- fmt.Sprintf("Writing: %s", text)
			// Simulate some work
			time.Sleep(time.Millisecond * 500)
		}
	}()
	return ch
}

func main() {
	chanel := write("Hello, World!")

	for i := 0; i < 5; i++ {
		fmt.Println(<-chanel)
	}
	fmt.Println("Done writing.")
}
