package db

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/gocarina/gocsv"
	"github.com/gofrs/uuid"
)

// Order represents a row of the csv
type Order struct { // Our example struct, you can use "-" to ignore a field
	ID       string    `csv:"ID"`
	ClientID string    `csv:"CLIENTID"`
	CarVIN   string    `csv:"CARVIN"`
	Quantity int64     `csv:"QUANTITY"`
	Date     time.Time `csv:"DATETIME"`
}

// OrderTable variable that has cars data loaded from csv
var OrderTable []*Order

// init loads in ram memory all data, this is execute when the module is imported.
func init() {
	ordersFile, err := os.OpenFile("db/orders.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	// it's necessary to verify if there is an error, in that case a panic must be raised.
	if err != nil {
		panic(err)
	}
	defer ordersFile.Close()
	OrderTable = []*Order{}
	if err := gocsv.UnmarshalFile(ordersFile, &OrderTable); err != nil { // Loads clients from file
		panic(err)
	}
}

// Save inserts a new order in csv
func (o *Order) Save() error {
	ordersFile, err := os.OpenFile("db/orders.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	if _, err := ordersFile.Seek(0, 0); err != nil { // Go to the start of the file
		panic(err)
	}
	id, err := uuid.NewV4()
	if err != nil {
		return err
	}
	o.ID = id.String()
	o.Date = time.Now().UTC()
	OrderTable = append(OrderTable, o)
	err = gocsv.MarshalFile(&OrderTable, ordersFile) // Use this to save the CSV back to the file
	return err
}

// FindOrderByID searches a client by id and returns the object if it is found
func FindOrderByID(id string) (*Order, error) {
	for _, order := range OrderTable {
		if order.ID == id {
			return order, nil
		}
	}
	return &Order{}, errors.New("Not found")
}

// FindAllOrdersByClientID searches all order by client and returns an slice with all the orders
func FindAllOrdersByClientID(id string) (Order, error) {
	// slice

	for _, order := range OrderTable {
		if order.ID == id {
			return *order, nil
		}
	}
	return Order{}, errors.New("Not found")
}

// PrintAllOrderTable displays all rows in the console.
func PrintAllOrderTable() {
	csvContent, err := gocsv.MarshalString(&OrderTable)
	if err != nil {
		panic(err)
	}
	fmt.Println(csvContent)
}
