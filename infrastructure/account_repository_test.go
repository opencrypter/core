package infrastructure_test

import (
	"github.com/opencrypter/api/domain"
	"github.com/opencrypter/api/infrastructure"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGormAccountRepository_AccountOfId(t *testing.T) {
	repository := infrastructure.NewAccountRepository()
	account := domain.NewAccount(uuid.NewV4().String(), uuid.NewV4().String(), uuid.NewV4().String(), "first", "api-key", "secret")
	currency := domain.NewCurrency(uuid.NewV4().String(), "Bitcoin", "BTC")
	balance := domain.NewBalance(uuid.NewV4().String(), account.ID, currency.ID, 10, false)
	balance.Currency = currency
	account.Balances = []domain.Balance{*balance}
	repository.Save(account)

	t.Run("It should return the account with its balances", func(t *testing.T) {
		result, err := repository.AccountOfId(account.ID)
		assert.NoError(t, err)
		assert.Equal(t, account, result)
	})
}

func TestGormAccountRepository_Save(t *testing.T) {
	repository := infrastructure.NewAccountRepository()
	account := domain.NewAccount(uuid.NewV4().String(), uuid.NewV4().String(), uuid.NewV4().String(), "first", "api-key", "secret")

	t.Run("It should save a new account", func(t *testing.T) {
		repository.Save(account)
		savedAccount, _ := repository.AccountOfId(account.ID)
		assert.NotNil(t, savedAccount)
	})

	t.Run("It should save an existing account", func(t *testing.T) {
		repository.Save(account)
		newName := "new-name"
		account.Name = &newName
		repository.Save(account)

		updatedAccount, err := repository.AccountOfId(account.ID)
		assert.NoError(t, err)
		assert.Equal(t, &newName, updatedAccount.Name)
	})

	t.Run("It should return an error on receive an invalid account", func(t *testing.T) {
		account := &domain.Account{ID: uuid.NewV4().String()}
		assert.Error(t, repository.Save(account))
		_, err := repository.AccountOfId(account.ID)
		assert.Error(t, err)
	})
}

func TestGormAccountRepository_AllOfDeviceId(t *testing.T) {
	repository := infrastructure.NewAccountRepository()
	deviceId := uuid.NewV4().String()
	repository.Save(domain.NewAccount(uuid.NewV4().String(), deviceId, uuid.NewV4().String(), "first", "api-key", "secret"))
	repository.Save(domain.NewAccount(uuid.NewV4().String(), deviceId, uuid.NewV4().String(), "second", "api-key", "secret"))

	t.Run("It should return all device accounts", func(t *testing.T) {
		devices := repository.AllOfDeviceId(deviceId)
		assert.Len(t, devices, 2)
	})

	t.Run("It should return nothing", func(t *testing.T) {
		devices := repository.AllOfDeviceId(uuid.NewV4().String())
		assert.Empty(t, devices)
	})
}
