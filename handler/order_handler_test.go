package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"weservefood/models"
	"weservefood/repository"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestPingServer(t *testing.T) {
	req, err := http.NewRequest("GET", "/ping", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(PingServer)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "Hello From the Server!!", rr.Body.String())
}

func TestPlaceOrder(t *testing.T) {
	order := models.Order{
		Email:   "test@example.com",
		Address: "123 Test St",
	}
	orderJSON, _ := json.Marshal(order)

	req, err := http.NewRequest("POST", "/place-order", bytes.NewBuffer(orderJSON))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(PlaceOrder)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var createdOrder models.Order
	err = json.NewDecoder(rr.Body).Decode(&createdOrder)
	assert.NoError(t, err)
	assert.Equal(t, order.Email, createdOrder.Email)
	assert.Equal(t, order.Address, createdOrder.Address)
}

func TestGetOrder(t *testing.T) {
	order := models.Order{
		Email:   "test@example.com",
		Address: "123 Test St",
	}
	_, _ = repository.CreateOrder(order)

	req, err := http.NewRequest("GET", "/get-order?email=test@example.com", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetOrder)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var orders []models.Order
	err = json.NewDecoder(rr.Body).Decode(&orders)
	assert.NoError(t, err)
	assert.NotEmpty(t, orders)
	assert.Equal(t, order.Email, orders[0].Email)
}

func TestGetAllOrders(t *testing.T) {
	order := models.Order{
		Email:   "test@example.com",
		Address: "123 Test St",
	}
	_, _ = repository.CreateOrder(order)

	req, err := http.NewRequest("GET", "/get-all-orders", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetAllOrders)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var orders []models.Order
	err = json.NewDecoder(rr.Body).Decode(&orders)
	assert.NoError(t, err)
	assert.NotEmpty(t, orders)
}

func TestCancelOrder(t *testing.T) {
	order := models.Order{
		Email:   "test@example.com",
		Address: "123 Test St",
	}
	createdOrder, _ := repository.CreateOrder(order)

	req, err := http.NewRequest("DELETE", "/cancel-order/test@example.com/"+createdOrder.ID, nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/cancel-order/{email}/{id}", CancelOrder).Methods("DELETE")
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var responseMessage string
	err = json.NewDecoder(rr.Body).Decode(&responseMessage)
	assert.NoError(t, err)
	assert.Equal(t, createdOrder.ID+" Order Cancelled Successfully", responseMessage)
}

func TestUpdateAddress(t *testing.T) {
	order := models.Order{
		Email:   "test@example.com",
		Address: "123 Test St",
	}
	createdOrder, _ := repository.CreateOrder(order)

	updateData := map[string]string{"new_address": "456 New St"}
	updateJSON, _ := json.Marshal(updateData)

	req, err := http.NewRequest("PUT", "/update-address/test@example.com/"+createdOrder.ID, bytes.NewBuffer(updateJSON))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/update-address/{email}/{id}", UpdateAddress).Methods("PUT")
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var updatedOrder models.Order
	err = json.NewDecoder(rr.Body).Decode(&updatedOrder)
	assert.NoError(t, err)
	assert.Equal(t, "456 New St", updatedOrder.Address)
}
