package repository

import (
	"errors"
	"fmt"
	"strconv"
	"time"
	"weservefood/models"

	"math/rand"
)

var store = models.InMemoryStore{
	Orders: make(map[string]models.Order),
}

// Generate a unique order ID using the current timestamp and a random number
func generateOrderID() string {
	return time.Now().Format("202402102150405") + strconv.Itoa(rand.Intn(100))
}

// CreateOrder creates a new order and returns the order details
func CreateOrder(newOrder models.Order) (models.Order, error) {
	newOrder.ID = generateOrderID()
	newOrder.DeliveryTime = time.Now().Add(30 * time.Minute).Format("15:01:09")

	store.Mutex.Lock()
	store.Orders[newOrder.ID] = newOrder
	store.Mutex.Unlock()

	return newOrder, nil

}

// GetOrderByEmail retrieves all orders for a given email
func GetOrderByEmail(email string) ([]models.Order, error) {
	var userOrders []models.Order

	store.Mutex.Lock()
	for _, order := range store.Orders {
		if order.Email == email {
			userOrders = append(userOrders, order)
		}
	}

	store.Mutex.Unlock()

	if len(userOrders) == 0 {
		return nil, errors.New("no orders found for the given email")
	}

	return userOrders, nil
}

// GetAllOrders retrieves all active orders
func GetAllOrders() ([]models.Order, error) {

	store.Mutex.Lock()
	if len(store.Orders) == 0 {
		store.Mutex.Unlock()
		return nil, errors.New("no active orders found")
	}

	orders := make([]models.Order, 0, len(store.Orders))

	for _, order := range store.Orders {
		orders = append(orders, order)
	}
	store.Mutex.Unlock()

	return orders, nil
}

// UpdateAddress updates the delivery address for a given order
func UpdateAddress(email, orderID, newAddress string) (models.Order, error) {
	store.Mutex.Lock()
	defer store.Mutex.Unlock()

	order, exist := store.Orders[orderID]
	if !exist {
		return models.Order{}, errors.New("order not found")
	}
	if order.Email != email {
		return models.Order{}, errors.New("email does not match")
	}

	order.Address = newAddress
	store.Orders[orderID] = order

	return order, nil
}

// CancelOrder cancels an order by order ID and email
func CancelOrder(email, orderID string) (string, error) {
	store.Mutex.Lock()
	order, exist := store.Orders[orderID]
	if exist && order.Email == email {
		delete(store.Orders, orderID)
		store.Mutex.Unlock()
		return fmt.Sprintf("%s Order Cancelled Successfully", orderID), nil
	} else {
		store.Mutex.Unlock()
		return "", errors.New("order not found")
	}
}
