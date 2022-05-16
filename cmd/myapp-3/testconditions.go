package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Return a number
func somenumber() int {
	return -7
}

// Initialize seed to generate random numbers
func initializeSeed() {
	sec := time.Now().Unix()
	rand.Seed(sec)
}

// Generate a random number
func generateRandom() int32 {
	return rand.Int31n(10)
}

// Return the region and continent for a given city
func location(city string) (string, string) {
	var region string
	var continent string
	switch city {
	case "Delhi", "Hyderabad", "Mumbai", "Chennai", "Kochi":
		region, continent = "India", "Asia"
	case "Lafayette", "Louisville", "Boulder":
		region, continent = "Colorado", "USA"
	case "Irvine", "Los Angeles", "San Diego":
		region, continent = "California", "USA"
	case "Lisboa", "Porto":
		region, continent = "Portugal", "Europe"
	case "Madrid", "Barcelona":
		region, continent = "Spain", "Europe"
	default:
		region, continent = "Unknown", "Unknown"
	}
	return region, continent
}

// Main function
func main() {
	initializeSeed()
	myNumber := generateRandom()
	fmt.Println("Random number: ", myNumber)

	if myNumber%2 == 0 {
		fmt.Println(myNumber, "is even")
	} else {
		fmt.Println(myNumber, "is uneven")
	}

	switch myNumber {
	case 0:
		fmt.Println("zero...")
	case 1:
		fmt.Println("one...")
	case 2:
		fmt.Println("two...")
	default:
		fmt.Println("no match...")
	}

	if myNumber := somenumber(); myNumber < 0 {
		fmt.Println(myNumber, "is negative")
	} else if myNumber < 10 {
		fmt.Println(myNumber, "has 1 digit")
	} else {
		fmt.Println(myNumber, "has multiple digits")
	}

	cityName := "Lisboa"
	region, continent := location(cityName)
	fmt.Printf("My city is %s, which is located in %s, %s\n", cityName, region, continent)
}
