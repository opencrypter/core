package application_test

import (
	"github.com/golang/mock/gomock"
	"github.com/opencrypter/api/application"
	"github.com/opencrypter/api/domain"
	"github.com/opencrypter/api/mock"
	"github.com/satori/go.uuid"
	"testing"
)

func TestGetAllAccounts_Execute(t *testing.T) {
	mockedRepository := mock.NewMockAccountRepository(gomock.NewController(t))
	service := application.NewGetAllAccounts(mockedRepository)

	t.Run("It should return all found accounts", func(t *testing.T) {
		deviceId := uuid.NewV4().String()
		expected := domain.NewAccount(
			uuid.NewV4().String(),
			deviceId,
			uuid.NewV4().String(),
			"test",
			"api-key",
			"api-secret",
		)
		mockedRepository.EXPECT().AllOfDeviceId(deviceId).Return([]domain.Account{*expected})
		service.Execute(deviceId)
	})
}
