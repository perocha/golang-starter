package main

import "fmt"

func main() {
	/*
		ARRAYS
	*/
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

	/*
		SLICES
	*/
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

	// 	lettersSlice = ["A", "B", "C", "D", "E"]
	//  index =          0    1    2    3    4
	//  position =       1    2    3    4    5
	lettersSlice := []string{"A", "B", "C", "D", "E"}
	fmt.Println("\nCreating copies of a Slice")
	fmt.Println("Before: ", lettersSlice)

	// slice1 = lettersSlice[index, last position] = ["A", "B"]
	slice1 := lettersSlice[0:2]
	// slice2 = lettersSlice[index, last position] = ["B", "C", "D"]
	slice2 := lettersSlice[1:4]

	// By changing slice2, it also changes the underlying lettersSlice slice, so all 3 slices will have this change
	slice2[1] = "Z"
	fmt.Println("After step 1 -->")
	fmt.Println("lettersSlice: ", lettersSlice)
	fmt.Println("Slice1: ", slice1)
	fmt.Println("Slice2: ", slice2)

	// Use the copy function to copy a Slice
	slice3 := make([]string, 3)
	copy(slice3, lettersSlice[1:4])

	// Changing again slice2, it will now change lettersSlice and slice2, but not slice3
	slice2[1] = "ZZ"
	fmt.Println("After step 2 -->")
	fmt.Println("lettersSlice: ", lettersSlice)
	fmt.Println("Slice1: ", slice1)
	fmt.Println("Slice2: ", slice2)
	fmt.Println("Slice3: ", slice3)

	/*
		MAPS
	*/
	// Map creation
	studentsAge := make(map[string]int)
	fmt.Println(studentsAge)
	studentsAge["Celia"] = 23
	studentsAge["Salomeja"] = 21

	// Print content of a map
	fmt.Println(studentsAge)
	fmt.Println("Print student age: ", studentsAge["Celia"])

	// Boolean for an element that doesn't exist in the map
	age, exist := studentsAge["Pedro"]
	if exist {
		fmt.Println("Student age is: ", age)
	} else {
		fmt.Println("Student not found")
	}

	// Add new elements to a map
	studentsAge["Pedro"] = 30
	studentsAge["Pedro"] = 31 // replace the value of an existing key
	studentsAge["Laura"] = 29
	fmt.Println("Add elements: ", studentsAge)

	// Delete elements from a map
	delete(studentsAge, "Pedro")
	fmt.Println("Remove elements: ", studentsAge)

	// Loop the map
	fmt.Println("List of all students with a for loop")
	for name, age := range studentsAge {
		fmt.Printf("%s is %v years old\n", name, age)
	}
}
