package main

import (
	"fmt"
	"net/http"
	"time"
)

// Function that execute a http GET to a given URL (it will be executed concurrently)
func checkAPI(api string, ch chan string) {
	_, err := http.Get(api)
	if err != nil {
		ch <- fmt.Sprintf("ERROR: %s is down!\n", api)
		return
	}

	ch <- fmt.Sprintf("SUCCESS: %s is up and running!\n", api)
}

// Function to send a message to a channel
func send(ch chan string, message string) {
	time.Sleep(2 * time.Second)
	ch <- message
}

func main() {
	start := time.Now()

	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com/",
		"https://api.somewhereintheinternet.com/",
		"https://graph.microsoft.com",
	}

	// Create unbuffered channel and call function checkAPI concurrently
	ch := make(chan string)
	for _, api := range apis {
		go checkAPI(api, ch)
	}

	// Wait for all info to be received in the channel
	for i := 0; i < len(apis); i++ {
		fmt.Print(<-ch)
	}

	// Print how long it took to execute the concurrent functions
	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())

	// Send 4 messages to a buffered channel
	start = time.Now()
	size := 10
	myChannel := make(chan string, size)
	go send(myChannel, "one")
	go send(myChannel, "two")
	go send(myChannel, "three")
	go send(myChannel, "four")

	for i := 0; i < 4; i++ {
		fmt.Println(<-myChannel)
	}

	// Once completed, each function takes 2 seconds to execute and total execution is ~2.00361 seconds
	fmt.Println("Done!")
	elapsed = time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())
}
