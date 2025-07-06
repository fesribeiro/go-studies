package main

import "fmt"

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func worker(tasks <-chan int, results chan<- int) {
	for n := range tasks {
		results <- fibonacci(n) // Calculate Fibonacci and send result to results channel
	}
}

func main() {
	tasks := make(chan int, 45) // Create a buffered channel to hold tasks
	results := make(chan int, 45) // Create a buffered channel to hold results

	go worker(tasks, results) // Start a worker goroutine
	go worker(tasks, results) // Start another worker goroutine
	go worker(tasks, results) // Start another worker goroutine
	go worker(tasks, results) // Start another worker goroutine
	// go worker(tasks, results) // Start a third worker goroutine

	for i := 0; i < 45; i++ { // Send tasks to the worker
		tasks <- i
	}

	close(tasks) // Close the tasks channel to signal no more tasks will be sent
	for i := 0; i < 45; i++ { // Collect results from the
		result := <-results // Receive result from results channel
		fmt.Println("Fibonacci of", i, "is", result) // Print the result
	}

}
