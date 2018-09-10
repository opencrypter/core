package domain_test

import (
	"github.com/opencrypter/core/domain"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewExchange(t *testing.T) {
	id := uuid.NewV4().String()
	name := "test"
	tag := "test"
	exchange := domain.NewExchange(id, name, tag)

	assert.Equal(t, id, exchange.ID)
	assert.Equal(t, &name, exchange.Name)
	assert.Equal(t, &tag, exchange.Tag)
}
