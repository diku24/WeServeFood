package models

import (
	"sync"
)

type Order struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Email        string   `json:"email"`
	Address      string   `json:"address"`
	Items        []string `json:"items"`
	DeliveryTime string   `json:"delivery_time"`
}

type InMemoryStore struct {
	Orders map[string]Order
	Mutex  sync.Mutex
}
