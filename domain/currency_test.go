package domain_test

import (
	"github.com/opencrypter/api/domain"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCurrency(t *testing.T) {
	id := uuid.NewV4().String()
	name := "Bitcoin"
	symbol := "BTC"
	currency := domain.NewCurrency(id, name, symbol)
	assert.Equal(t, id, currency.ID)
	assert.Equal(t, &name, currency.Name)
	assert.Equal(t, &symbol, currency.Symbol)
}
