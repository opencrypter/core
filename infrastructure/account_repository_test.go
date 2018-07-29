package infrastructure

import (
	"github.com/opencrypter/api/domain"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGormAccountRepository_Save(t *testing.T) {
	repository := NewAccountRepository()

	name := "test"
	exchangeId := uuid.NewV4().String()
	apiKey := "api-key"
	apiSecret := "api-secret"

	t.Run("It should save a new account", func(t *testing.T) {
		account := &domain.Account{
			ID:         uuid.NewV4().String(),
			Name:       &name,
			ExchangeId: &exchangeId,
			ApiKey:     &apiKey,
			ApiSecret:  &apiSecret,
		}

		repository.Save(account)

		savedAccount, _ := repository.AccountOfId(account.ID)
		assert.NotNil(t, savedAccount)
	})

	t.Run("It should save an existing account", func(t *testing.T) {
		account := &domain.Account{
			ID:         uuid.NewV4().String(),
			Name:       &name,
			ExchangeId: &exchangeId,
			ApiKey:     &apiKey,
			ApiSecret:  &apiSecret,
		}

		repository.Save(account)

		newName := "new-name"
		account.Name = &newName
		repository.Save(account)

		updatedAccount, _ := repository.AccountOfId(account.ID)
		assert.Equal(t, &newName, updatedAccount.Name)
	})

	t.Run("It should return an error on receive an invalid account", func(t *testing.T) {
		account := &domain.Account{
			ID: uuid.NewV4().String(),
		}

		assert.Error(t, repository.Save(account))
		_, err := repository.AccountOfId(account.ID)
		assert.Error(t, err)
	})
}
