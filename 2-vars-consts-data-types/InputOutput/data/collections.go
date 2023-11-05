package data

import "fmt"

// Arrays
var Countries [10]string

// Slices (Dynamic Arrays)
var Slice []int

// Dictionaries/Maps
var Codes map[int]string

func init() {
	Countries[0] = "Argentina"
	Countries[1] = "Chile"
	Countries[2] = "Brazil"
	Countries[3] = "Peru"
	Countries[4] = "Columbia"
	Countries[5] = "Mexico"
	Countries[6] = "USA"
	Countries[7] = "Canada"
	Countries[8] = "Cuba"
	Countries[9] = "UK"

	qty := len(Countries)

	fmt.Println(qty)

	fmt.Println("Countries Saved!")
}