package main

import (
	"fmt"
)

var name = "Frontend Masters"

func init() {
	// fmt.Println("B")
}

// func calculateTax(price float32) (float32, float32) {
// 	return price * 0.09, price * 0.02
// }

// func calculateTaxWithName(price float32) (stateTax float32, cityTax float32) {
// 	stateTax = price * 0.09
// 	cityTax = price * 0.02
// 	return
// }

func birthday(pointerAge *int) {
	fmt.Printf("The pointer is %v and the value is %v\n", pointerAge, *pointerAge)

	*pointerAge++
}

func main() {
	// _, cityTax := calculateTax(100)

	// fmt.Println(cityTax)

	age := 22
	birthday(&age)
	fmt.Println(age)

	// PrintData()
	// fmt.Println(data.MaxSpeed)
}