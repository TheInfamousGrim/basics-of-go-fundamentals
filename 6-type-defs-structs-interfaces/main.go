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

	kyle := data.NewInstructor("Kyle", "Simpson")

	goCourse := data.Course{Id: 2, Name: "Go Fundamentals", Instructor: max}
	
	fmt.Println(max.Print())
	fmt.Println(kyle.Print())
	fmt.Println(goCourse.String())

	// Swift workshop for embedding
	swiftWS := data.NewWorkshop("Swift with iOS", max)

	fmt.Printf("%v", swiftWS)
}