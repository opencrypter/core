package ui

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetExchangeTickers(t *testing.T) {
	router := NewRouter()

	t.Run("It should return not implemented status code", func(t *testing.T) {
		responseRecorder := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", "/exchanges/100cfe0b-78be-42c2-ba42-95d1f2c67336/tickers", nil)
		router.ServeHTTP(responseRecorder, request)
		assert.Equal(t, http.StatusNotImplemented, responseRecorder.Code)
	})
}

func TestGetTickerAlerts(t *testing.T) {
	router := NewRouter()

	t.Run("It should return not implemented status code", func(t *testing.T) {
		responseRecorder := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", "/tickers/100cfe0b-78be-42c2-ba42-95d1f2c67336/alerts", nil)
		router.ServeHTTP(responseRecorder, request)
		assert.Equal(t, http.StatusNotImplemented, responseRecorder.Code)
	})
}

func TestPutTickerAlert(t *testing.T) {
	router := NewRouter()

	t.Run("It should return not implemented status code", func(t *testing.T) {
		path := "/tickers/100cfe0b-78be-42c2-ba42-95d1f2c67336/alerts/100cfe0b-78be-42c2-ba42-95d1f2c67336"
		responseRecorder := httptest.NewRecorder()
		request, _ := http.NewRequest("PUT", path, nil)
		router.ServeHTTP(responseRecorder, request)
		assert.Equal(t, http.StatusNotImplemented, responseRecorder.Code)
	})
}

func TestDeleteTickerAlert(t *testing.T) {
	router := NewRouter()

	t.Run("It should return not implemented status code", func(t *testing.T) {
		path := "/tickers/100cfe0b-78be-42c2-ba42-95d1f2c67336/alerts/100cfe0b-78be-42c2-ba42-95d1f2c67336"
		responseRecorder := httptest.NewRecorder()
		request, _ := http.NewRequest("DELETE", path, nil)
		router.ServeHTTP(responseRecorder, request)
		assert.Equal(t, http.StatusNotImplemented, responseRecorder.Code)
	})
}
