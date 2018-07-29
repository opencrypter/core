package infrastructure_test

import (
	"fmt"
	"github.com/opencrypter/api/domain"
	"github.com/opencrypter/api/infrastructure"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGormAccountRepository_Save(t *testing.T) {
	repository := infrastructure.NewAccountRepository()

	name := "test"
	exchangeId := uuid.NewV4().String()
	deviceId := uuid.NewV4().String()
	apiKey := "api-key"
	apiSecret := "api-secret"
	account := domain.NewAccount(uuid.NewV4().String(), deviceId, exchangeId, name, apiKey, apiSecret)

	t.Run("It should save a new account", func(t *testing.T) {
		repository.Save(account)

		savedAccount, _ := repository.AccountOfId(account.ID)
		assert.NotNil(t, savedAccount)
	})

	t.Run("It should save an existing account", func(t *testing.T) {
		err := repository.Save(account)
		fmt.Println(err)
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
	repository.Save(
		domain.NewAccount(uuid.NewV4().String(), deviceId, uuid.NewV4().String(), "first", "api-key", "secret"),
	)
	repository.Save(
		domain.NewAccount(uuid.NewV4().String(), deviceId, uuid.NewV4().String(), "second", "api-key", "secret"),
	)

	t.Run("It should return all device accounts", func(t *testing.T) {
		devices := repository.AllOfDeviceId(deviceId)
		assert.Len(t, devices, 2)
	})

	t.Run("It should return nothing", func(t *testing.T) {
		devices := repository.AllOfDeviceId(uuid.NewV4().String())
		assert.Empty(t, devices)
	})
}
