package models

import "time"

// Order represents a order object
type Order struct { // Our example struct, you can use "-" to ignore a field
	ID       string    `json:"id"`
	ClientID string    `json:"client_id"`
	CarVIN   string    `json:"vin"`
	Quantity int64     `json:"quantity"`
	Date     time.Time `json:"-"`
}
