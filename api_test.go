// +build integration

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNotImplementedEndpoints(t *testing.T) {
	router := NewRouter()

	t.Run("Test post account", func(t *testing.T) {
		assertNotImplemented(t, router, "POST", "/devices")
	})

	t.Run("Test put account", func(t *testing.T) {
		assertNotImplemented(t, router, "PUT", "/accounts/100cfe0b-78be-42c2-ba42-95d1f2c67336")
	})

	t.Run("Test get account", func(t *testing.T) {
		assertNotImplemented(t, router, "GET", "/accounts/100cfe0b-78be-42c2-ba42-95d1f2c67336")
	})

	t.Run("Test get balances", func(t *testing.T) {
		assertNotImplemented(t, router, "GET", "/accounts/100cfe0b-78be-42c2-ba42-95d1f2c67336/balances")
	})

	t.Run("Test get exchanges", func(t *testing.T) {
		assertNotImplemented(t, router, "GET", "/exchanges")
	})

	t.Run("Test get exchange detail", func(t *testing.T) {
		assertNotImplemented(t, router, "GET", "/exchanges/100cfe0b-78be-42c2-ba42-95d1f2c67336")
	})

	t.Run("Test get exchange tickers", func(t *testing.T) {
		assertNotImplemented(t, router, "GET", "/exchanges/100cfe0b-78be-42c2-ba42-95d1f2c67336/tickers")
	})

	t.Run("Test get ticker alerts", func(t *testing.T) {
		assertNotImplemented(t, router, "GET", "/tickers/100cfe0b-78be-42c2-ba42-95d1f2c67336/alerts")
	})

	t.Run("Test put ticker alert", func(t *testing.T) {
		path := "/tickers/100cfe0b-78be-42c2-ba42-95d1f2c67336/alerts/100cfe0b-78be-42c2-ba42-95d1f2c67336"
		assertNotImplemented(t, router, "PUT", path)
	})

	t.Run("Test delete ticker alert", func(t *testing.T) {
		path := "/tickers/100cfe0b-78be-42c2-ba42-95d1f2c67336/alerts/100cfe0b-78be-42c2-ba42-95d1f2c67336"
		assertNotImplemented(t, router, "DELETE", path)
	})
}

func assertNotImplemented(t *testing.T, router *gin.Engine, method string, path string) {
	responseRecorder := httptest.NewRecorder()
	request, _ := http.NewRequest("PUT", "/accounts/100cfe0b-78be-42c2-ba42-95d1f2c67336", nil)
	router.ServeHTTP(responseRecorder, request)
	assert.Equal(t, http.StatusNotImplemented, responseRecorder.Code)
}
