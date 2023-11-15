package data

// Type alias
type integer = int
type json = map[string]string

// New types
type distance float64 // miles
type distanceKm float64 // km

// var x integer

// Add methods for types

// Method
func (miles distance) ToKm() distanceKm {
	return distanceKm(1.60934 * miles)
}

func (kilometres distanceKm) ToMiles() distance {
	return distance(kilometres / 1.60934)
}

func Test() {
	d := distance(4.5)
	km := d.ToKm()
	returnedD := km.ToMiles()
	print(d)
	print(returnedD)
}
