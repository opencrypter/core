package infrastructure_test

import (
	"github.com/opencrypter/core/domain"
	"github.com/opencrypter/core/infrastructure"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGormExchangeRepository_Save(t *testing.T) {
	repository := infrastructure.NewExchangeRepository()

	t.Run("It should save the exchange", func(t *testing.T) {
		exchange := domain.NewExchange(uuid.NewV4().String(), "test-save-exchange", "save-exchange")
		repository.Save(exchange)
		savedExchange, err := repository.ExchangeOfId(exchange.ID)
		assert.NoError(t, err)
		assert.Equal(t, exchange, savedExchange)
	})

	t.Run("It should return an error on invalid exchange", func(t *testing.T) {
		exchange := domain.NewExchange(uuid.NewV4().String(), "test-save-exchange", "save-exchange")
		exchange.Name = nil
		repository.Save(exchange)
		_, err := repository.ExchangeOfId(exchange.ID)
		assert.Error(t, err)
	})
}

func TestGormExchangeRepository_All(t *testing.T) {
	repository := infrastructure.NewExchangeRepository()

	t.Run("It should return all exchanges", func(t *testing.T) {
		assert.NotEmpty(t, repository.All())
	})
}
