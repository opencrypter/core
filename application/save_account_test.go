package application_test

import (
	"github.com/golang/mock/gomock"
	"github.com/opencrypter/api/application"
	"github.com/opencrypter/api/domain"
	"github.com/opencrypter/api/mock"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
	"testing"
)

func TestSaveAccount_Execute(t *testing.T) {
	mockedRepository := mock.NewMockAccountRepository(gomock.NewController(t))
	service := application.NewSaveAccount(mockedRepository)

	exchangeId := uuid.NewV4().String()
	deviceId := uuid.NewV4().String()
	name := "test"
	ApiKey := "api-key"
	ApiSecret := "api-secret"

	expected := &domain.Account{
		ID:         uuid.NewV4().String(),
		DeviceId:   &deviceId,
		ExchangeId: &exchangeId,
		Name:       &name,
		ApiKey:     &ApiKey,
		ApiSecret:  &ApiSecret,
	}

	t.Run("It should save the account", func(t *testing.T) {
		mockedRepository.EXPECT().Save(expected).Return(nil)
		service.Execute(expected.ID, &deviceId, &exchangeId, &name, &ApiKey, &ApiSecret)
	})

	t.Run("It should return an error on repository fail", func(t *testing.T) {
		mockedRepository.EXPECT().Save(expected).Return(errors.New("error"))
		service.Execute(expected.ID, &deviceId, &exchangeId, &name, &ApiKey, &ApiSecret)
	})
}
