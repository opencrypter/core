package domain_test

import (
	"github.com/opencrypter/core/domain"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewAccount(t *testing.T) {
	id := uuid.NewV4().String()
	name := "test"
	exchangeId := uuid.NewV4().String()
	deviceId := uuid.NewV4().String()
	apiKey := "api-key"
	apiSecret := "api-secret"
	account := domain.NewAccount(id, deviceId, exchangeId, name, apiKey, apiSecret)

	assert.Equal(t, id, account.ID)
	assert.Equal(t, &name, account.Name)
	assert.Equal(t, &deviceId, account.DeviceId)
	assert.Equal(t, &exchangeId, account.ExchangeId)
	assert.Equal(t, &apiKey, account.ApiKey)
	assert.Equal(t, &apiSecret, account.ApiSecret)
}

func TestNewBalance(t *testing.T) {
	id := uuid.NewV4().String()
	accountId := uuid.NewV4().String()
	currencyId := uuid.NewV4().String()
	volume := 100.5
	hasAlert := true

	balance := domain.NewBalance(id, accountId, currencyId, volume, hasAlert)
	assert.Equal(t, id, balance.ID)
	assert.Equal(t, &accountId, balance.AccountID)
	assert.Equal(t, &currencyId, balance.CurrencyID)
	assert.Equal(t, &volume, balance.Volume)
	assert.Equal(t, &hasAlert, balance.HasAlert)
}
