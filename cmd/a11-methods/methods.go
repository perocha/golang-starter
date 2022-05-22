package main

import (
	"fmt"
	"math"
	"strings"
)

/****************************************************
	STRUCTS
****************************************************/
// Basic structs
type triangle struct {
	size int
}

type square struct {
	size int
}

type circle struct {
	radius int
}

// Embed struct
type coloredTriangle struct {
	triangle
	color string
}

// Custom string type
type myTypeString string

/****************************************************
	METHODS
****************************************************/
// Return the perimeter of the triangle
func (t triangle) perimeter() int {
	return t.size * 3
}

// Return the perimeter of the square
func (s square) perimeter() int {
	return s.size * 4
}

// Return the perimeter of the circle (using math.Pi)
func (c circle) perimeter() float64 {
	return float64(2) * math.Pi * float64(c.radius)
}

// Overload method
func (t coloredTriangle) perimeter() int {
	return t.size * 3 * 2
}

// Double the size of the triangle (using pointer)
func (t *triangle) doubleSize() {
	t.size *= 2
}

// Return an uppercase string
func (s myTypeString) Upper() string {
	return strings.ToUpper(string(s))
}

/****************************************************
	MAIN
****************************************************/
func main() {
	// Basic methods
	t := triangle{3}
	s := square{4}
	c := circle{5}
	fmt.Println("Perimeter (triangle):", t.perimeter())
	fmt.Println("Perimeter (square):", s.perimeter())
	fmt.Println("Perimeter (circle):", c.perimeter())

	// Update the variable using a pointer
	t.doubleSize()
	fmt.Println("Size of triangle:", t.size)
	fmt.Println("Perimeter (double size triangle):", t.perimeter())

	// Custom type string
	varString := "this is just a test"
	mySpecialString := myTypeString(varString)
	fmt.Println("Basic type string: ", varString)
	fmt.Println("Custom string type: mySpecialString = ", mySpecialString)
	fmt.Println("Custom string type: mySpecialString.Upper() = ", mySpecialString.Upper())

	// Use of an embed method
	newTriangle := coloredTriangle{t, "Blue"}
	fmt.Println("newTriangle.color: ", newTriangle.color)
	fmt.Println("newTriangle.size: ", newTriangle.size)
	fmt.Println("newTriangle.perimeter: ", newTriangle.triangle.perimeter())
	// Use overload perimeter method
	fmt.Println("newTriangle.overload perimeter: ", newTriangle.perimeter())

}
