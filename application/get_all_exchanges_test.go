package application_test

import (
	"github.com/golang/mock/gomock"
	"github.com/opencrypter/core/application"
	"github.com/opencrypter/core/domain"
	"github.com/opencrypter/core/mock"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAllExchanges_Execute(t *testing.T) {
	mockedRepository := mock.NewMockExchangeRepository(gomock.NewController(t))
	service := application.NewGetAllExchanges(mockedRepository)

	t.Run("It should return all exchanges", func(t *testing.T) {
		expected := []domain.Exchange{*domain.NewExchange(uuid.NewV4().String(), "test-service", "test")}
		mockedRepository.EXPECT().All().Return(expected)
		assert.Equal(t, expected, service.Execute())
	})
}
