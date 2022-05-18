package main

import "fmt"

func main() {
	var a [3]int
	a[1] = 10
	fmt.Println(a[0])
	fmt.Println(a[1])
	fmt.Println(a[len(a)-1])

	cities := [5]string{"Lisbon", "Madrid", "Barcelona", "Vilnius", "Budapest"}
	fmt.Println("My fav cities: ", cities)

	moreCities := [...]string{"Lisbon", "Madrid", "Barcelona", "Vilnius", "Budapest", "Seattle", "New York", "Rome"}
	fmt.Println("My fav cities: ", moreCities)
	fmt.Println("Size: ", len(moreCities))

	numbers := [...]int{5: 23, 99: -1}
	fmt.Println("Position 0: ", numbers[0])
	fmt.Println("Position 1: ", numbers[1])
	fmt.Println("Position 5: ", numbers[5])
	fmt.Println("Position 99: ", numbers[99])
	fmt.Println("Len: ", len(numbers))

	multiDimensionalArray()
}

// Create a multi dimensional array
func multiDimensionalArray() {
	fmt.Println("\nmultiDimensionalArray")
	var threeD [3][5][2]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			for k := 0; k < 2; k++ {
				threeD[i][j][k] = (i + 1) * (j + 1) * (k + 1)
			}
		}
		fmt.Println("Row", i, threeD[i])
	}
	fmt.Println("All at once:", threeD)
}
