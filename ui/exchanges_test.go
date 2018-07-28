package ui_test

import (
	"github.com/opencrypter/api/ui"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetExchanges(t *testing.T) {
	router := ui.NewRouter()

	t.Run("It should return not implemented status code", func(t *testing.T) {
		responseRecorder := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", "/exchanges", nil)
		router.ServeHTTP(responseRecorder, request)
		assert.Equal(t, http.StatusNotImplemented, responseRecorder.Code)
	})
}

func TestGetExchangeDetail(t *testing.T) {
	router := ui.NewRouter()

	t.Run("It should return not implemented status code", func(t *testing.T) {
		responseRecorder := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", "/exchanges/100cfe0b-78be-42c2-ba42-95d1f2c67336", nil)
		router.ServeHTTP(responseRecorder, request)
		assert.Equal(t, http.StatusNotImplemented, responseRecorder.Code)
	})
}
