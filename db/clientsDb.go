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
	// verifica que el cliente tenga id y nombre para poder guardarlo
	if c.ID == "" || c.Name == "" {
		// Como mi validación genera el error, se crea un error
		return errors.New("Properties ID and Name should not be empty")
	}
	clientsFile, err := os.OpenFile("db/clients.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	if _, err := clientsFile.Seek(0, 0); err != nil { // Go to the start of the file
		errors.Wrap(err, "Clients Save")
	}
	// usa una libreria para generar un GUID
	id, err := uuid.NewV4()
	if err != nil {
		errors.Wrap(err, "Clients Save")
	}
	// convierte el tipo de dato  UUID del paquete uuid a una cadena
	c.ID = id.String()
	// inserta al cliente al slice
	ClientTable = append(ClientTable, *c)
	err = gocsv.MarshalFile(&ClientTable, clientsFile) // Use this to save the CSV back to the file
	if err != nil {
		// Como el error puede venir de una libreria externa se usa wrap
		return errors.Wrap(err, "Clients Save")
	}
	return nil
}

// FindClientByID searches a client by id and returns the object if it is found
func FindClientByID(id string) (Client, error) {
	// Hace un for range para ir obteniendo cada uno de los clientes en la tabla
	for i, client := range ClientTable {
		// verifica si existe el id
		if client.ID == id {
			// retorna el objeto con los datos y el id especificado
			return client, nil
		}
		log.Println(i, client)
	}
	// retorna un error. Como mi validación genera el error, se crea un error con error.New
	return Client{}, errors.New("Not Found")
}

// FindClients returns all clients in table
func FindClients() ([]Client, error) {
	// Verifica que haya datos en la tabla (slice)
	if len(ClientTable) > 0 {
		return ClientTable, nil
	}
	// devuelve un error si el slice esta vacio. Como mi validación genera el error, se crea un error con error.New
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
