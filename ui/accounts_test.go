package ui_test

import (
	"bytes"
	"encoding/json"
	"github.com/opencrypter/api/application"
	"github.com/opencrypter/api/domain"
	"github.com/opencrypter/api/infrastructure"
	"github.com/opencrypter/api/ui"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAllAccounts(t *testing.T) {
	router := ui.NewRouter()
	account := domain.NewAccount(
		uuid.NewV4().String(),
		suite.existingDevice.ID,
		uuid.NewV4().String(),
		"test-accounts",
		"api-key-accounts",
		"api-secret-accounts",
	)
	infrastructure.NewAccountRepository().Save(account)

	t.Run("It should return existing accounts", func(t *testing.T) {
		expected := []application.SecureAccountDto{{account.ID, account.ExchangeId, account.Name}}
		expectedJson, _ := json.Marshal(expected)

		responseRecorder := httptest.NewRecorder()
		request := suite.newAuthenticatedRequest(requestData{method: "GET", path: "/accounts"})
		router.ServeHTTP(responseRecorder, request)
		assert.Equal(t, http.StatusOK, responseRecorder.Code)
		assert.JSONEq(t, string(expectedJson), responseRecorder.Body.String())
	})

	t.Run("It should return forbidden on missing credentials", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", "/accounts", nil)
		router.ServeHTTP(recorder, request)
		assert.Equal(t, http.StatusForbidden, recorder.Code)
	})
}

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

		request := suite.newAuthenticatedRequest(requestData{
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
		recorder := httptest.NewRecorder()
		buffer := new(bytes.Buffer)
		json.NewEncoder(buffer).Encode(&ui.AccountDto{Id: "invalid"})

		request := suite.newAuthenticatedRequest(requestData{
			buffer: buffer,
			method: "PUT",
			path:   "/accounts/invalid",
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
		request := suite.newAuthenticatedRequest(requestData{
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
		request := suite.newAuthenticatedRequest(requestData{
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
