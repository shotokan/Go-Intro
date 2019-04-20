package db

import (
	"fmt"
	"log"
	"os"

	"github.com/gocarina/gocsv"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
)

// Client represents a row of the csv
type Client struct { // Our example struct, you can use "-" to ignore a field
	ID       string `csv:"ID"`
	Name     string `csv:"NAME"`
	LastName string `csv:"LASTNAME"`
	Address  string `csv:"ADDRESS"`
}

// ClientTable variable that has cars data loaded from csv
var ClientTable []Client

// init loads in ram memory all data, this is execute when the module is imported.
func init() {
	clientsFile, err := os.OpenFile("db/clients.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer clientsFile.Close()
	ClientTable = []Client{}
	if err := gocsv.UnmarshalFile(clientsFile, &ClientTable); err != nil { // Load clients from file
		panic(err)
	}
}

// Save inserts a new client in csv
func (c *Client) Save() error {
	clientsFile, err := os.OpenFile("db/clients.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	if _, err := clientsFile.Seek(0, 0); err != nil { // Go to the start of the file
		errors.Wrap(err, "Clients Save")
	}
	id, err := uuid.NewV4()
	if err != nil {
		errors.Wrap(err, "Clients Save")
	}
	c.ID = id.String()
	ClientTable = append(ClientTable, *c)
	err = gocsv.MarshalFile(&ClientTable, clientsFile) // Use this to save the CSV back to the file
	if err != nil {
		return errors.Wrap(err, "Clients Save")
	}
	return nil
}

// FindClientByID searches a client by id and returns the object if it is found
func FindClientByID(id string) (Client, error) {
	for i, client := range ClientTable {
		if client.ID == id {
			return client, nil
		}
		log.Println(i, client)
	}
	return Client{}, errors.Wrap(errors.New("Not Found"), "FindClientByID")
}

// FindClients returns all clients in table
func FindClients() ([]Client, error) {
	if len(ClientTable) > 0 {
		return ClientTable, nil
	}
	return nil, errors.New("Empty")
}

// PrintAllClientTable displays all rows in the console.
func PrintAllClientTable() {
	csvContent, err := gocsv.MarshalString(&ClientTable)
	if err != nil {
		errors.Wrap(err, "Clients PrintAllClientTable")
	}
	fmt.Println(csvContent)
}
