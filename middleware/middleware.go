package middleware

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// LoggingMiddleware logs the incoming requests
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		log.Printf("%s - %s %s %s", req.Method, req.Host, req.RequestURI, req.RemoteAddr)
		next.ServeHTTP(rw, req)
	})
}

// ValidationMiddleware validates the incoming requests
func ValidationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodPost:
			if !validatePostRequest(rw, req) {
				return
			}
		case http.MethodGet:
			if !validateGetRequest(rw, req) {
				return
			}
		case http.MethodPut:
			if !validatePutRequest(rw, req) {
				return
			}
		case http.MethodDelete:
			if !validateDeleteRequest(rw, req) {
				return
			}
		default:

			http.Error(rw, "Invalid Request Method", http.StatusMethodNotAllowed)
			return
		}
		next.ServeHTTP(rw, req)
	})
}

// validatePostRequest validates the POST request
func validatePostRequest(rw http.ResponseWriter, req *http.Request) bool {
	if req.Body == nil {
		log.Println(" Validation Failed: Missing request body")
		http.Error(rw, " Validation Failed: Missing request body", http.StatusBadRequest)
		return false
	}
	return true
}

// validateGetRequest validates the GET request
func validateGetRequest(rw http.ResponseWriter, req *http.Request) bool {
	email := req.URL.Query().Get("email")
	if email == "" && req.URL.Path == "/get-order" {
		log.Println("  Validation Failed: Missing email in query parameter")
		http.Error(rw, " Validation Failed: Missing email in query parameter", http.StatusBadRequest)
		return false
	}
	return true
}

// validatePutRequest validates the PUT request
func validatePutRequest(rw http.ResponseWriter, req *http.Request) bool {
	vars := mux.Vars(req)
	email := vars["email"]
	orderId := vars["id"]
	if email == "" || orderId == "" {
		log.Println("  Validation Failed: Missing email or orderID parameter")
		http.Error(rw, " Validation Failed: Missing email or orderID parameter", http.StatusBadRequest)
		return false
	}
	return true
}

// validateDeleteRequest validates the DELETE request
func validateDeleteRequest(rw http.ResponseWriter, req *http.Request) bool {
	vars := mux.Vars(req)
	email := vars["email"]
	orderId := vars["id"]
	if email == "" || orderId == "" {
		log.Println(" Validation Failed: Missing email or orderID Parameter")
		http.Error(rw, " Validation Failed: Missing email or orderID  parameter", http.StatusBadRequest)
		return false
	}
	return true
}
