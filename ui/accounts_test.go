package ui

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPutAccount(t *testing.T) {
	router := NewRouter()

	t.Run("It should return not implemented status code", func(t *testing.T) {
		responseRecorder := httptest.NewRecorder()
		request, _ := http.NewRequest("PUT", "/accounts/100cfe0b-78be-42c2-ba42-95d1f2c67336", nil)
		router.ServeHTTP(responseRecorder, request)
		assert.Equal(t, http.StatusNotImplemented, responseRecorder.Code)
	})
}

func TestGetAccount(t *testing.T) {
	router := NewRouter()

	t.Run("It should return not implemented status code", func(t *testing.T) {
		responseRecorder := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", "/accounts/100cfe0b-78be-42c2-ba42-95d1f2c67336", nil)
		router.ServeHTTP(responseRecorder, request)
		assert.Equal(t, http.StatusNotImplemented, responseRecorder.Code)
	})
}

func TestGetBalances(t *testing.T) {
	router := NewRouter()

	t.Run("It should return not implemented status code", func(t *testing.T) {
		responseRecorder := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", "/accounts/100cfe0b-78be-42c2-ba42-95d1f2c67336/balances", nil)
		router.ServeHTTP(responseRecorder, request)
		assert.Equal(t, http.StatusNotImplemented, responseRecorder.Code)
	})
}
