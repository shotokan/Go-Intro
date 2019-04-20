package controllers

import (
	"raw-service/db"
	"raw-service/models"

	"github.com/pkg/errors"
)

// GetClients returns a list of all clients
func GetClients() ([]models.Client, error) {
	// inicializa un objeto de tipo Client
	clientsResp := []models.Client{}
	// ejecuta la función FindClients la cual devuelve un listado de los clientes de tipo Client del paquete db
	clients, err := db.FindClients()
	if err != nil {
		return clientsResp, errors.Wrap(err, "Clients")
	}
	// Como se devuelve un slice de tipo models.Client (Client del paquete models), se debe parsear los datos del paquete db.Client a models.Client
	for _, client := range clients {
		c := models.Client{}
		c.ID = client.ID
		c.Address = client.Address
		c.Name = client.Name
		c.LastName = client.Name
		clientsResp = append(clientsResp, c)
	}
	// se retorna el nuevo slice con los datos convertidos a models.Client y un nil ya que no hubo error
	return clientsResp, nil
}

// GetClient searches a client by their ID and returns their data.
func GetClient(id string) (models.Client, error) {
	// llama a FindClientByID del paquete db y le pasa un parámetro de tipo string para poder encontrar al cliente
	// dicha función devuelve un objeto de tipo db.Client o un error
	client, err := db.FindClientByID(id)
	if err != nil {
		// Hace el wrap del error para mandarlo a la función que la llamó
		return models.Client{}, errors.Wrap(err, "GetClient error")
	}
	// Se convierte el objeto db.Client a models.Client, ya que el método devuelve este último tipo
	model := models.Client{
		ID:       client.ID,
		Address:  client.Address,
		Name:     client.Name,
		LastName: client.LastName,
	}
	// se retorna el cliente de topo models.Client y un nil ya que no hubo error
	return model, nil
}
