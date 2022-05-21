package main

import "fmt"

func main() {
	// Very simple array
	var a [3]int
	a[1] = 10
	fmt.Println(a[0])
	fmt.Println(a[1])
	fmt.Println(a[len(a)-1])

	// Create a simple array of strings
	cities := [5]string{"Lisbon", "Madrid", "Barcelona", "Vilnius", "Budapest"}
	fmt.Println("My fav cities: ", cities)

	// Create an array of strings without specifying size
	moreCities := [...]string{"Lisbon", "Madrid", "Barcelona", "Vilnius", "Budapest", "Seattle", "New York", "Rome"}
	fmt.Println("My fav cities: ", moreCities)
	fmt.Println("Size: ", len(moreCities))

	// Create an array of integer, specifying 2 values
	numbers := [...]int{5: 23, 99: -1}
	fmt.Println("Position 0: ", numbers[0])
	fmt.Println("Position 1: ", numbers[1])
	fmt.Println("Position 5: ", numbers[5])
	fmt.Println("Position 99: ", numbers[99])
	fmt.Println("Len: ", len(numbers))

	// Create a multi dimensional array
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

	// Slice
	monthSlice := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	fmt.Println("\nTest using Slice")
	fmt.Println(monthSlice)
	fmt.Println("Length:", len(monthSlice))
	fmt.Println("Capacity:", cap(monthSlice))

	// Create additional slices from the original Slice
	quarter1 := monthSlice[0:3]
	quarter2 := monthSlice[3:6]
	quarter3 := monthSlice[6:9]
	quarter4 := monthSlice[9:12]
	fmt.Println("\nPrint parts of the Slice, the length and capacity:")
	fmt.Println(quarter1, len(quarter1), cap(quarter1))
	fmt.Println(quarter2, len(quarter2), cap(quarter2))
	fmt.Println(quarter3, len(quarter3), cap(quarter3))
	fmt.Println(quarter4, len(quarter4), cap(quarter4))

	// Append elements to an empty Slice
	var numberSlice []int
	fmt.Println("\nAppending elements to a Slice")
	for i := 0; i < 10; i++ {
		numberSlice = append(numberSlice, i)
		fmt.Printf("%d\tCapacity = %d\t%v\n", i, cap(numberSlice), numberSlice)
	}

}
