package main

import "fmt"

func main() {
	var operation string
	var number1, number2 int

	fmt.Println("CALCULATOR GO 1.0")
	fmt.Println("=================")

	// Operation Input
	fmt.Println("Which operation do you want to perform? (add, subtract, multiply, divide)")
	fmt.Scanf("%s", operation)

	// First input
	fmt.Println("Enter first number")
	fmt.Scanf("%d", &number1)	

	// First input
	fmt.Println("Enter second number")
	fmt.Scanf("%d", &number2)	

	// Switch statement for calculator operations
	switch operation {
	case "add":
		fmt.Println(number1 + number2)
	case "subtract":
		fmt.Println(number1 - number2)
	case "multiply":
		fmt.Println(number1 * number2)
	case "divide":
		fmt.Println(number1 / number2)
	}

}
