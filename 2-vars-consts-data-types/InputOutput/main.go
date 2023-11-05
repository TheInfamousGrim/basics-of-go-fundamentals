package main

import (
	"fmt"

	"frontendmasters.com/go/io/data"
)

// Global variable
// var url = "https://frontendmasters.com"
var name = "Frontend Masters"

func main() {
	// const pi float32 = 3.12
	// const maxSpeed byte = 60

	// Function scoped variables
	// message := "Hello from Go"
	// price := 34.4

	// code block scoped variables

	// print(message, price, url)
	
	PrintData()
	fmt.Println(data.MaxSpeed)
}