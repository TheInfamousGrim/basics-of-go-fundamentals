package main

import (
	"fmt"

	"frontendmasters.com/go/server/data"
)

func main() {
	max := data.Instructor {
		Id: 3, FirstName: "Maximiliano",
	}

	max.LastName = "Fabiano"
	max.Score = 96
	
	fmt.Println(max.Print())
}