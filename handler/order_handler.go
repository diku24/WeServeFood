package handler

import (
	"encoding/json"
	"net/http"
	"weservefood/models"
	"weservefood/repository"

	"github.com/gorilla/mux"
)

const ContentTypeHeader string = "Content-Type"
const ApplicationJson string = "application/json"

// @Summary Ping Server
// @Description Check server availability
// @Produce plain
// @Success 200 {string} string "Hello From the Server!!"
// @Router /ping [get]
func PingServer(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("Hello From the Server!!"))
}

// @Summary Place an order
// @Description Create a new food order
// @Accept json
// @Produce json
// @Param order body models.Order true "Order Details"
// @Success 200 {object} models.Order "Order Details"
// @Failure 400 {string} string "Invalid Request Payload"
// @Failure 500 {string} string "Internal Server Error"
// @Router /place-order [post]
func PlaceOrder(rw http.ResponseWriter, req *http.Request) {
	var newOrder models.Order

	if err := json.NewDecoder(req.Body).Decode(&newOrder); err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	order, err := repository.CreateOrder(newOrder)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(rw).Encode(order); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}

// @Summary Get user orders
// @Description Retrieve all orders for a given email
// @Produce json
// @Param email query string true "User Email"
// @Success 200 {object} models.Order
// @Failure 404 {string} string "Order not found"
// @Router /get-order [get]
func GetOrder(rw http.ResponseWriter, req *http.Request) {
	email := req.URL.Query().Get("email")

	order, err := repository.GetOrderByEmail(email)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusNotFound)
		return
	}

	rw.Header().Set(ContentTypeHeader, ApplicationJson)
	//json.NewEncoder(rw).Encode(order)
	if err := json.NewEncoder(rw).Encode(order); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

}

// @Summary Get all orders
// @Description Retrieve all active orders
// @Produce json
// @Success 200 {array} []models.Order
// @Failure 404 {string} string "No active orders found"
// @Router /get-all-orders [get]
func GetAllOrders(rw http.ResponseWriter, req *http.Request) {

	orders, err := repository.GetAllOrders()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusNotFound)
		return
	}

	rw.Header().Set(ContentTypeHeader, ApplicationJson)
	//json.NewEncoder(rw).Encode(orders)
	if err := json.NewEncoder(rw).Encode(orders); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}

// @Summary Cancel an order
// @Description Cancel an order by order ID and email
// @Produce json
// @Param email path string true "User Email"
// @Param id path string true "Order ID"
// @Success 200 {string} string "Order Cancelled Successfully"
// @Router /cancel-order/{email}/{id} [delete]
func CancelOrder(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	email := vars["email"]
	orderID := vars["id"]

	message, err := repository.CancelOrder(email, orderID)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusNotFound)
	}

	rw.Header().Set(ContentTypeHeader, ApplicationJson)
	//json.NewEncoder(rw).Encode(message)
	if err := json.NewEncoder(rw).Encode(message); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

}

// @Summary Update address
// @Description Update the delivery address for an order
// @Produce json
// @Param email path string true "User Email"
// @Param id path string true "Order ID"
// @Param new_address query string true "New Address"
// @Success 200 {object} models.Order
// @Failure 400 {string} string "unable to update new address"
// @Router /update-address/{email}/{id} [put]
func UpdateAddress(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	email := vars["email"]
	orderID := vars["id"]

	var requestData struct {
		NewAddress string `json:"new_address"`
	}

	if err := json.NewDecoder(req.Body).Decode(&requestData); err != nil {
		http.Error(rw, "unable to update new address", http.StatusBadRequest)
		return
	}

	updatedOrder, err := repository.UpdateAddress(email, orderID, requestData.NewAddress)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	rw.Header().Set(ContentTypeHeader, ApplicationJson)
	//json.NewEncoder(rw).Encode(updatedOrder)
	if err := json.NewEncoder(rw).Encode(updatedOrder); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

}
