package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Function to calculate Fibonacci
func fib(number float64, ch chan string) {
	x, y := 1.0, 1.0
	for i := 0; i < int(number); i++ {
		x, y = y, x+y
	}

	r := rand.Intn(3)
	time.Sleep(time.Duration(r) * time.Second)

	ch <- fmt.Sprintf("Fib(%v): %v\n", number, x)
}

func main() {
	start := time.Now()

	// Create a buffered channel (with size)
	size := 15
	myChannel := make(chan string, size)
	fmt.Printf("Channel capacity: %d\n", cap(myChannel))

	// For loop to execute all Fibonacci calculations in parallel
	for i := 0; i < size; i++ {
		go fib(float64(i), myChannel)
	}

	// Get all the responses from the Fibonacci functions
	for i := 0; i < size; i++ {
		fibonacciResult := <-myChannel
		fmt.Println(fibonacciResult)
	}

	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())
}
