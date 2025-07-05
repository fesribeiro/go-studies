package main

import (
	"fmt"
	"sync"
	"time"
)

func write(text string) {
	for i:= 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(1 * time.Second)
	}
}

func main() {

	var waitGroup sync.WaitGroup
	waitGroup.Add(2) // We have two goroutines to wait for

	go func() {
		write("Hello, World!")
		waitGroup.Done() // Signal that this goroutine is done
	}()

	go func() {
		write("Hello, Go!")
		waitGroup.Done() // Signal that this goroutine is done
	}()

	waitGroup.Wait() // Wait for all goroutines to finish
	fmt.Println("All goroutines have finished executing.")
}