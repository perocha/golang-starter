package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

// Error variables
var ErrEmptyString = errors.New("invalid input::empty string")
var ErrStringTooBig = errors.New("invalid input::string is too big")

func main() {
	// Create a log file
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.SetOutput(file)
	log.SetPrefix("main(): ")
	log.Print("Hey, I'm a new log file!")

	inputString := "Name"
	result, err := testFunction(inputString)
	if err == nil {
		fmt.Println("Result call 1: ", result)
	} else {
		log.Printf("Error call 1: %v", err)
		fmt.Println("Err call 1: ", err)
	}

	inputString = ""
	result, err = testFunction(inputString)
	if err == nil {
		fmt.Println("Result call 2: ", result)
	} else {
		log.Printf("Error call 2: %v", err)
		fmt.Println("Err call 2: ", err)
	}

	inputString = "This is just a test"
	result, err = testFunction(inputString)
	if err == nil {
		fmt.Println("Result call 3: ", result)
	} else {
		log.Printf("Error call 3: %v", err)
		fmt.Println("Err call 3: ", err)
		log.Panic("Can't continue execution!!")
		//log.Fatal("Can't continue execution!!")
		fmt.Println("This message is not printed, program stops with panic")
	}
}

func testFunction(input string) (int, error) {
	stringLenght := len(input)

	if stringLenght == 0 {
		return stringLenght, ErrEmptyString
	}

	if stringLenght > 10 {
		return stringLenght, ErrStringTooBig
	}

	return stringLenght, nil
}
