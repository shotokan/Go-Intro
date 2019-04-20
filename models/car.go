package models

// Car represents a car object
type Car struct {
	VIN   string `csv:"VIN"`
	Model string `csv:"MODEL"`
	Color string `csv:"COLOR"`
	Make  string `csv:"MAKE"`
	Year  int    `csv:"YEAR"`
}
