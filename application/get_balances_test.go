package application_test

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/opencrypter/core/application"
	"github.com/opencrypter/core/domain"
	"github.com/opencrypter/core/mock"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetBalances_Execute(t *testing.T) {
	mockedRepository := mock.NewMockAccountRepository(gomock.NewController(t))
	service := application.NewGetBalances(mockedRepository)

	t.Run("It should return all balances", func(t *testing.T) {
		account := domain.NewAccount(
			uuid.NewV4().String(),
			uuid.NewV4().String(),
			uuid.NewV4().String(),
			"test",
			"test",
			"test",
		)
		balance := domain.NewBalance(
			uuid.NewV4().String(),
			account.ID,
			uuid.NewV4().String(),
			10,
			false,
		)
		account.Balances = []domain.Balance{*balance}
		mockedRepository.EXPECT().AccountOfId(account.ID).Return(account, nil)
		balances, err := service.Execute(account.ID)
		assert.NoError(t, err)
		assert.Equal(t, account.Balances, balances)
	})

	t.Run("It should return error on missing account", func(t *testing.T) {
		accountId := uuid.NewV4().String()
		mockedRepository.EXPECT().AccountOfId(accountId).Return(nil, errors.New("error"))
		balances, err := service.Execute(accountId)
		assert.Error(t, err)
		assert.Nil(t, balances)
	})
}
