package db

import (
	"os"

	"github.com/gocarina/gocsv"
	"github.com/pkg/errors"
)

// Cars represents a row of the csv
type Car struct { // Our example struct, you can use "-" to ignore a field
	VIN   string `csv:"VIN"`
	Model string `csv:"MODEL"`
	Color string `csv:"COLOR"`
	Make  string `csv:"MAKE"`
	Year  int    `csv:"YEAR"`
}

// CarsTable variable that has cars data loaded from csv
var CarsTable []*Car

// init loads in ram memory all data, this is execute when the module is imported.
func init() {
	carsFile, err := os.OpenFile("db/cars.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer carsFile.Close()
	CarsTable = []*Car{}
	if err := gocsv.UnmarshalFile(carsFile, &CarsTable); err != nil { // Load clients from file
		panic(err)
	}
}

// Save new cars added into csv
func (c *Car) Add() error {
	carsFile, err := os.OpenFile("db/cars.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		errors.Wrap(err, "Car Add")
	}
	if _, err := carsFile.Seek(0, 0); err != nil { // Go to the start of the file
		errors.Wrap(err, "Car Add")
	}
	err = gocsv.MarshalFile(&CarsTable, carsFile) // Use this to save the CSV back to the file
	return errors.Wrap(err, "Car Add")
}

// Save new cars added into csv
func (c *Car) Save() error {
	carsFile, err := os.OpenFile("db/cars.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		errors.Wrap(err, "Car Save")
	}
	if _, err := carsFile.Seek(0, 0); err != nil { // Go to the start of the file
		errors.Wrap(err, "Car Save")
	}
	err = gocsv.MarshalFile(&CarsTable, carsFile) // Use this to save the CSV back to the file
	return err
}
