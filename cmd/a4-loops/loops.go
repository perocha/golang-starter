package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Call a function with a basic for loop
	minValue := 1
	maxValue := 101
	sumResult := basicForLoop(minValue, maxValue)
	fmt.Printf("Sum of %v..%v is %v\n", minValue, maxValue, sumResult)

	// Infinite loop
	infiniteForLoop()

	// Function using defer
	usingDefer()

	// Print prime numbers < 20
	fmt.Println("Prime numbers less than 20:")

	for number := 0; number <= 20; number++ {
		if findprimes(number) {
			fmt.Printf("%v ", number)
		}
	}
}

// Function to test a basic for loop
func basicForLoop(minValue int, maxValue int) int {
	ResultSum := 0
	for i := minValue; i <= maxValue; i++ {
		ResultSum += i
	}
	return ResultSum
}

// Function to test an infinite loop
func infiniteForLoop() {
	var num int32
	sec := time.Now().Unix()
	rand.Seed(sec)

	for {
		fmt.Print("Writing inside the loop...")
		if num = rand.Int31n(10); num == 5 {
			fmt.Println("finish!")
			break
		}
		fmt.Println(num)
	}
}

// For using defer statement
func usingDefer() {
	for i := 1; i <= 4; i++ {
		defer fmt.Println("deferred", -i)
		fmt.Println("regular", i)
	}
}

// Function to find the prime numbers
func findprimes(number int) bool {
	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}

	if number > 1 {
		return true
	} else {
		return false
	}
}
