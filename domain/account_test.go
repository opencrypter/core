package domain_test

import (
	"github.com/opencrypter/api/domain"
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
