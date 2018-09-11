package main_test

import (
	"github.com/gin-gonic/gin/json"
	"github.com/opencrypter/core/domain"
	"github.com/opencrypter/core/infrastructure"
	"github.com/opencrypter/core/ui/http"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetExchanges(t *testing.T) {
	router := main.NewRouter()
	database := infrastructure.Database
	database.Delete(domain.Exchange{})

	t.Run("It should return exchanges", func(t *testing.T) {
		exchanges := []domain.Exchange{*domain.NewExchange(uuid.NewV4().String(), "test", "test")}
		expected, _ := json.Marshal(exchanges)
		infrastructure.NewExchangeRepository().Save(&exchanges[0])

		recorder := httptest.NewRecorder()
		request := suite.newAuthenticatedRequest(requestData{method: "GET", path: "/exchanges"})
		router.ServeHTTP(recorder, request)
		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.JSONEq(t, string(expected), recorder.Body.String())
	})

	t.Run("It should return forbidden on missing credentials", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", "/exchanges", nil)
		router.ServeHTTP(recorder, request)
		assert.Equal(t, http.StatusForbidden, recorder.Code)
	})
}

func TestGetExchangeDetail(t *testing.T) {
	router := main.NewRouter()

	t.Run("It should return not implemented status code", func(t *testing.T) {
		responseRecorder := httptest.NewRecorder()
		request := suite.newAuthenticatedRequest(requestData{
			method: "GET",
			path:   "/exchanges/100cfe0b-78be-42c2-ba42-95d1f2c67336",
		})
		router.ServeHTTP(responseRecorder, request)
		assert.Equal(t, http.StatusNotImplemented, responseRecorder.Code)
	})

	t.Run("It should return forbidden on missing credentials", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", "/exchanges/100cfe0b-78be-42c2-ba42-95d1f2c67336", nil)
		router.ServeHTTP(recorder, request)
		assert.Equal(t, http.StatusForbidden, recorder.Code)
	})
}
