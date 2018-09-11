package main_test

import (
	"github.com/opencrypter/core/ui/http"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetExchangeTickers(t *testing.T) {
	router := main.NewRouter()

	t.Run("It should return not implemented status code", func(t *testing.T) {
		responseRecorder := httptest.NewRecorder()
		request := suite.newAuthenticatedRequest(requestData{
			method: "GET",
			path:   "/exchanges/100cfe0b-78be-42c2-ba42-95d1f2c67336/tickers",
		})
		router.ServeHTTP(responseRecorder, request)
		assert.Equal(t, http.StatusNotImplemented, responseRecorder.Code)
	})
}

func TestGetTickerAlerts(t *testing.T) {
	router := main.NewRouter()

	t.Run("It should return not implemented status code", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		request := suite.newAuthenticatedRequest(requestData{
			method: "GET",
			path:   "/tickers/100cfe0b-78be-42c2-ba42-95d1f2c67336/alerts",
		})
		router.ServeHTTP(recorder, request)
		assert.Equal(t, http.StatusNotImplemented, recorder.Code)
	})
}

func TestPutTickerAlert(t *testing.T) {
	router := main.NewRouter()

	t.Run("It should return not implemented status code", func(t *testing.T) {
		path := "/tickers/100cfe0b-78be-42c2-ba42-95d1f2c67336/alerts/100cfe0b-78be-42c2-ba42-95d1f2c67336"
		responseRecorder := httptest.NewRecorder()
		request, _ := http.NewRequest("PUT", path, nil)
		router.ServeHTTP(responseRecorder, request)
		assert.Equal(t, http.StatusNotImplemented, responseRecorder.Code)
	})
}

func TestDeleteTickerAlert(t *testing.T) {
	router := main.NewRouter()

	t.Run("It should return not implemented status code", func(t *testing.T) {
		path := "/tickers/100cfe0b-78be-42c2-ba42-95d1f2c67336/alerts/100cfe0b-78be-42c2-ba42-95d1f2c67336"
		responseRecorder := httptest.NewRecorder()
		request, _ := http.NewRequest("DELETE", path, nil)
		router.ServeHTTP(responseRecorder, request)
		assert.Equal(t, http.StatusNotImplemented, responseRecorder.Code)
	})
}
