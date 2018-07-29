package ui_test

import (
	"bytes"
	"encoding/json"
	"github.com/opencrypter/api/infrastructure"
	"github.com/opencrypter/api/ui"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPutAccount(t *testing.T) {
	router := ui.NewRouter()

	id := uuid.NewV4().String()
	exchangeId := uuid.NewV4().String()
	name := "account-test"
	apiKey := "apiKey"
	apiSecret := "apiSecret"

	t.Run("It should save the account", func(t *testing.T) {
		dto := ui.AccountDto{Id: id, ExchangeId: &exchangeId, Name: &name, ApiKey: &apiKey, ApiSecret: &apiSecret}
		recorder := httptest.NewRecorder()
		buffer := new(bytes.Buffer)
		json.NewEncoder(buffer).Encode(&dto)

		request := createAuthenticatedRequest(requestData{
			device: suite.existingDevice,
			buffer: buffer,
			method: "PUT",
			path:   "/accounts/" + dto.Id,
		})
		router.ServeHTTP(recorder, request)

		savedAccount, _ := infrastructure.NewAccountRepository().AccountOfId(dto.Id)
		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.NotNil(t, savedAccount)
	})

	t.Run("It should return bad request on invalid account", func(t *testing.T) {
		dto := ui.AccountDto{Id: "invalid"}
		recorder := httptest.NewRecorder()
		buffer := new(bytes.Buffer)
		json.NewEncoder(buffer).Encode(&dto)

		request := createAuthenticatedRequest(requestData{
			device: suite.existingDevice,
			buffer: buffer,
			method: "PUT",
			path:   "/accounts/" + dto.Id,
		})
		router.ServeHTTP(recorder, request)
		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("It should return forbidden on missing credentials", func(t *testing.T) {
		dto := ui.AccountDto{Id: id, ExchangeId: &exchangeId, Name: &name, ApiKey: &apiKey, ApiSecret: &apiSecret}
		recorder := httptest.NewRecorder()
		buffer := new(bytes.Buffer)
		json.NewEncoder(buffer).Encode(&dto)
		request, _ := http.NewRequest("PUT", "/accounts/"+dto.Id, buffer)
		router.ServeHTTP(recorder, request)
		assert.Equal(t, http.StatusForbidden, recorder.Code)
	})
}

func TestGetAccount(t *testing.T) {
	router := ui.NewRouter()

	t.Run("It should return not implemented status code", func(t *testing.T) {
		responseRecorder := httptest.NewRecorder()
		request := createAuthenticatedRequest(requestData{
			device: suite.existingDevice,
			buffer: new(bytes.Buffer),
			method: "GET",
			path:   "/accounts/100cfe0b-78be-42c2-ba42-95d1f2c67336",
		})
		router.ServeHTTP(responseRecorder, request)
		assert.Equal(t, http.StatusNotImplemented, responseRecorder.Code)
	})

	t.Run("It should return forbidden on missing credentials", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", "/accounts/100cfe0b-78be-42c2-ba42-95d1f2c67336", nil)
		router.ServeHTTP(recorder, request)
		assert.Equal(t, http.StatusForbidden, recorder.Code)
	})
}

func TestGetBalances(t *testing.T) {
	router := ui.NewRouter()

	t.Run("It should return not implemented status code", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		request := createAuthenticatedRequest(requestData{
			device: suite.existingDevice,
			buffer: new(bytes.Buffer),
			method: "GET",
			path:   "/accounts/100cfe0b-78be-42c2-ba42-95d1f2c67336/balances",
		})
		router.ServeHTTP(recorder, request)
		assert.Equal(t, http.StatusNotImplemented, recorder.Code)
	})

	t.Run("It should return forbidden on missing credentials", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", "/accounts/100cfe0b-78be-42c2-ba42-95d1f2c67336/balances", nil)
		router.ServeHTTP(recorder, request)
		assert.Equal(t, http.StatusForbidden, recorder.Code)
	})
}
