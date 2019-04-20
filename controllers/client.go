package controllers

import (
	"raw-service/db"
	"raw-service/models"

	"github.com/pkg/errors"
)

// GetClients returns a list of all clients
func GetClients() ([]models.Client, error) {
	clientsResp := []models.Client{}
	clients, err := db.FindClients()
	if err != nil {
		return clientsResp, errors.Wrap(err, "Clients")
	}
	for _, client := range clients {
		c := models.Client{}
		c.ID = client.ID
		c.Address = client.Address
		c.Name = client.Name
		c.LastName = client.Name
		clientsResp = append(clientsResp, c)
	}
	return clientsResp, nil
}

// GetClient searches a client by their ID and returns their data.
func GetClient(id string) (models.Client, error) {
	client, err := db.FindClientByID(id)
	if err != nil {
		return models.Client{}, errors.Wrap(err, "GetClient error")
	}
	model := models.Client{
		ID:       client.ID,
		Address:  client.Address,
		Name:     client.Name,
		LastName: client.LastName,
	}
	return model, nil
}
