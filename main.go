package main

import (
	"log"
	"net/http"
	"weservefood/handler"
	"weservefood/middleware"

	_ "weservefood/docs"

	"github.com/gorilla/mux"
	swagger "github.com/swaggo/http-swagger"
)

// @title WeServeFood Delivery Order Management API
// @version 1.0
// @description API for managing food delivery orders
// @host localhost:8383
// @BasePath /
func main() {
	route := mux.NewRouter()

	route.Use(middleware.LoggingMiddleware)
	route.Use(middleware.ValidationMiddleware)

	route.HandleFunc("/ping", handler.PingServer).Methods("GET")
	route.HandleFunc("/place-order", handler.PlaceOrder).Methods("POST")
	route.HandleFunc("/get-order", handler.GetOrder).Methods("GET")
	route.HandleFunc("/get-all-orders", handler.GetAllOrders).Methods("GET")
	route.HandleFunc("/cancel-order/{email}/{id}", handler.CancelOrder).Methods("DELETE")
	route.HandleFunc("/update-address/{email}/{id}", handler.UpdateAddress).Methods("PUT")

	route.PathPrefix("/swagger/").Handler(swagger.Handler()).Methods(http.MethodGet)

	log.Println("Starting server on port 8383")
	http.ListenAndServe(":8383", route)
}
