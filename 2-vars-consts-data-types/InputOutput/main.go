package main

// Global variable
var url = "https://frontendmasters.com"

func main() {
	// const pi float32 = 3.12
	// const maxSpeed byte = 60

	// Function scoped variables
	message := "Hello from Go"
	price := 34.4

	// code block scoped variables

	print(message, price, url)
}