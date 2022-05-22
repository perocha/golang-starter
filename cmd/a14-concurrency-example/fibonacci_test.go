package main

import "testing"

func TestFibonacci(t *testing.T) {
	// Create a channel with size 1
	size := 1
	myChannel := make(chan string, size)

	// For loop to execute all Fibonacci calculations in parallel
	sequence := 13
	go fib(float64(sequence), myChannel)

	// Fibonacci result must be 377 with the specific string format
	fibonacciResult := <-myChannel
	if fibonacciResult != "Fib(13): 377\n" {
		t.Errorf("incorrect value of Fibonacci sequence: $$%s$$", fibonacciResult)
	}
}
