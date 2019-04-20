package models

// Client represents a client object
type Client struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"lastname"`
	Address  string `json:"address"`
}
