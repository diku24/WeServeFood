package middleware

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestLoggingMiddleware(t *testing.T) {
	handler := LoggingMiddleware(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
	}))

	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestValidationMiddlewarePostRequest(t *testing.T) {
	handler := ValidationMiddleware(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
	}))

	req, _ := http.NewRequest(http.MethodPost, "/", strings.NewReader("test body"))
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestValidationMiddlewarePostRequestNoBody(t *testing.T) {
	handler := ValidationMiddleware(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
	}))

	req, _ := http.NewRequest(http.MethodPost, "/", nil)
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestValidationMiddlewareGetRequest(t *testing.T) {
	handler := ValidationMiddleware(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
	}))

	req, _ := http.NewRequest(http.MethodGet, "/get-order?email=test@example.com", nil)
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestValidationMiddlewareGetRequestNoEmail(t *testing.T) {
	handler := ValidationMiddleware(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
	}))

	req, _ := http.NewRequest(http.MethodGet, "/get-order", nil)
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestValidationMiddlewarePutRequest(t *testing.T) {
	handler := ValidationMiddleware(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
	}))

	req, _ := http.NewRequest(http.MethodPut, "/update-order/test@example.com/123", nil)
	req = mux.SetURLVars(req, map[string]string{"email": "test@example.com", "id": "123"})
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestValidationMiddlewarePutRequestMissingParams(t *testing.T) {
	handler := ValidationMiddleware(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
	}))

	req, _ := http.NewRequest(http.MethodPut, "/update-order/test@example.com/", nil)
	req = mux.SetURLVars(req, map[string]string{"email": "test@example.com"})
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestValidationMiddlewareDeleteRequest(t *testing.T) {
	handler := ValidationMiddleware(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
	}))

	req, _ := http.NewRequest(http.MethodDelete, "/delete-order/test@example.com/123", nil)
	req = mux.SetURLVars(req, map[string]string{"email": "test@example.com", "id": "123"})
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestValidationMiddlewareDeleteRequestMissingParams(t *testing.T) {
	handler := ValidationMiddleware(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
	}))

	req, _ := http.NewRequest(http.MethodDelete, "/delete-order/test@example.com/", nil)
	req = mux.SetURLVars(req, map[string]string{"email": "test@example.com"})
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}
