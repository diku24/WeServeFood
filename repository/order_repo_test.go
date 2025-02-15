package repository

import (
	"testing"
	"weservefood/models"

	"github.com/stretchr/testify/assert"
)

func TestCreateOrder(t *testing.T) {
	newOrder := models.Order{
		Email:   "test@example.com",
		Address: "123 Test St",
	}

	createdOrder, err := CreateOrder(newOrder)
	assert.NoError(t, err)
	assert.NotEmpty(t, createdOrder.ID)
	assert.Equal(t, newOrder.Email, createdOrder.Email)
	assert.Equal(t, newOrder.Address, createdOrder.Address)
}

func TestGetOrderByEmail(t *testing.T) {
	email := "test@example.com"
	newOrder := models.Order{
		Email:   email,
		Address: "123 Test St",
	}

	_, err := CreateOrder(newOrder)
	assert.NoError(t, err)

	orders, err := GetOrderByEmail(email)
	assert.NoError(t, err)
	assert.NotEmpty(t, orders)
	assert.Equal(t, email, orders[0].Email)
}

func TestGetOrderByEmailNoOrders(t *testing.T) {
	email := "noorders@example.com"
	orders, err := GetOrderByEmail(email)
	assert.Error(t, err)
	assert.Nil(t, orders)
}

func TestGetAllOrders(t *testing.T) {
	newOrder := models.Order{
		Email:   "test@example.com",
		Address: "123 Test St",
	}

	_, err := CreateOrder(newOrder)
	assert.NoError(t, err)

	orders, err := GetAllOrders()
	assert.NoError(t, err)
	assert.NotEmpty(t, orders)
}

func TestGetAllOrdersNoOrders(t *testing.T) {
	store.Orders = make(map[string]models.Order) // Clear the store
	orders, err := GetAllOrders()
	assert.Error(t, err)
	assert.Nil(t, orders)
}

func TestUpdateAddress(t *testing.T) {
	email := "test@example.com"
	newOrder := models.Order{
		Email:   email,
		Address: "123 Test St",
	}

	createdOrder, err := CreateOrder(newOrder)
	assert.NoError(t, err)

	newAddress := "456 New St"
	updatedOrder, err := UpdateAddress(email, createdOrder.ID, newAddress)
	assert.NoError(t, err)
	assert.Equal(t, newAddress, updatedOrder.Address)
}

func TestUpdateAddressOrderNotFound(t *testing.T) {
	email := "test@example.com"
	newAddress := "456 New St"
	_, err := UpdateAddress(email, "nonexistentID", newAddress)
	assert.Error(t, err)
}

func TestUpdateAddressEmailMismatch(t *testing.T) {
	email := "test@example.com"
	newOrder := models.Order{
		Email:   email,
		Address: "123 Test St",
	}

	createdOrder, err := CreateOrder(newOrder)
	assert.NoError(t, err)

	newAddress := "456 New St"
	_, err = UpdateAddress("wrong@example.com", createdOrder.ID, newAddress)
	assert.Error(t, err)
}

func TestCancelOrder(t *testing.T) {
	email := "test@example.com"
	newOrder := models.Order{
		Email:   email,
		Address: "123 Test St",
	}

	createdOrder, err := CreateOrder(newOrder)
	assert.NoError(t, err)

	msg, err := CancelOrder(email, createdOrder.ID)
	assert.NoError(t, err)
	assert.Contains(t, msg, "Order Cancelled Successfully")
}

func TestCancelOrderNotFound(t *testing.T) {
	email := "test@example.com"
	_, err := CancelOrder(email, "nonexistentID")
	assert.Error(t, err)
}
