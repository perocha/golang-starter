package main

import (
	"fmt"
	"math"
)

// Interface declaration
type Shape interface {
	Perimeter() float64
	Area() float64
}

// Square struct declaration
type Square struct {
	size float64
}

func (s Square) Area() float64 {
	return s.size * s.size
}

func (s Square) Perimeter() float64 {
	return s.size * 4
}

// Circle struct declaration
type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// Function that accepts as an argument interface Shape to print its content
func printInformation(s Shape) {
	fmt.Printf("%T\n", s)
	fmt.Println("Area: ", s.Area())
	fmt.Println("Perimeter:", s.Perimeter())
	fmt.Println()
}

func main() {
	// Explicit creation of a var of type interface
	var s Shape = Square{3}
	// Create an instance of struct Circle, which will also be accepted by printInformation function
	c := Circle{5}

	// Call same function with similar types
	printInformation(s)
	printInformation(c)
}
