package main

import (
	"fmt"

	"function-control-error.com/go/io/data"
)

var name = "Frontend Masters"

func init() {
	// fmt.Println("B")
}

func calculateTax(price float32) (float32, float32) {
	return price * 0.09, price * 0.02
}

// func calculateTaxWithName(price float32) (stateTax float32, cityTax float32) {
// 	stateTax = price * 0.09
// 	cityTax = price * 0.02
// 	return
// }

func main() {
	_, cityTax := calculateTax(100)

	fmt.Println(cityTax)

	// PrintData()
	fmt.Println(data.MaxSpeed)
}