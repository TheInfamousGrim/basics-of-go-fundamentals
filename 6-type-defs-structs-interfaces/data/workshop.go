package data

import "time"

// Embeds properties
type Workshop struct {
	Course
	Instructor
	Date   time.Time
}

func NewWorkshop (name string, instructor Instructor) Workshop {
	w := Workshop {}

	w.Name = name
	// w.Course.Name = "an example"
	w.Instructor = instructor

	return w
}