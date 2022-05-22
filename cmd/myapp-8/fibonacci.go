package main

import "fmt"

/*
	The Fibonacci sequence is a list of numbers where each number is the sum of the previous Fibonacci numbers.
	For instance, the sequence of numbers for 6 is 1,1,2,3,5,8, for 7 is 1,1,2,3,5,8,13, for 8 is 1,1,2,3,5,8,13,21, and so on.
*/
func main() {
	var num int

	fmt.Print("What's the Fibonacci sequence you want? ")
	fmt.Scanln(&num)

	// Example 1
	fibonacci(1, num, 0, 1)

	// Example 2
	result := fibonacciArray(num)
	fmt.Println("Array result: ", result)
}

// Recursive function call
func fibonacci(currentSequence int, maxSequence int, previousValue int, currentValue int) {
	if currentSequence < maxSequence+1 {
		fmt.Printf("Seq: %v Value: %v\n", currentSequence, currentValue)
		newValue := previousValue + currentValue
		currentSequence++
		fibonacci(currentSequence, maxSequence, currentValue, newValue)
	}
}

// Using a for loop and array
func fibonacciArray(n int) []int {
	if n < 2 {
		return make([]int, 0)
	}

	nums := make([]int, n)
	nums[0], nums[1] = 1, 1

	for i := 2; i < n; i++ {
		nums[i] = nums[i-1] + nums[i-2]
	}

	return nums
}
