package application_test

import (
	"github.com/golang/mock/gomock"
	"github.com/opencrypter/api/application"
	"github.com/opencrypter/api/domain"
	"github.com/opencrypter/api/mock"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAllAccounts_Execute(t *testing.T) {
	mockedRepository := mock.NewMockAccountRepository(gomock.NewController(t))
	service := application.NewGetAllAccounts(mockedRepository)

	t.Run("It should return all found accounts", func(t *testing.T) {
		deviceId := uuid.NewV4().String()
		expectedFromDb := domain.NewAccount(
			uuid.NewV4().String(),
			deviceId,
			uuid.NewV4().String(),
			"test",
			"api-key",
			"api-secret",
		)
		expected := application.SecureAccountDto{
			Id:         expectedFromDb.ID,
			ExchangeId: expectedFromDb.ExchangeId,
			Name:       expectedFromDb.Name,
		}
		mockedRepository.EXPECT().AllOfDeviceId(deviceId).Return([]domain.Account{*expectedFromDb})
		accounts := service.Execute(deviceId)
		assert.Equal(t, []application.SecureAccountDto{expected}, accounts)
	})
}
