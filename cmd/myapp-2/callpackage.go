package main

import (
	"fmt"
	"log"

	"golang-starter/internal/internalapp1"
)

func main() {
	log.Printf("callpackage::start")

	returnResult := internalapp1.Sum(5, 6)
	fmt.Println(returnResult)

	fmt.Println(internalapp1.Test())

	log.Printf("callpackage::finish")
}
