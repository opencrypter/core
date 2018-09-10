package domain_test

import (
	"github.com/opencrypter/core/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDevice(t *testing.T) {
	id := "cfe64055-5059-4464-8c02-6f5500cbaf08"
	os := "ios"
	senderId := "abc"

	t.Run("It should create a device", func(t *testing.T) {
		device, _ := domain.NewDevice(id, os, &senderId)
		assert.Equal(t, id, device.ID)
		assert.Equal(t, &os, device.Os)
		assert.Equal(t, &senderId, device.SenderId)
		assert.NotEmpty(t, device.Secret)
	})

	t.Run("It should work without sender id", func(t *testing.T) {
		device, _ := domain.NewDevice(id, os, nil)
		assert.Nil(t, device.SenderId)
	})

	t.Run("It should return error on invalid id", func(t *testing.T) {
		_, err := domain.NewDevice("invalid", os, &senderId)
		assert.Error(t, err)
	})

	t.Run("It should return error on empty os", func(t *testing.T) {
		_, err := domain.NewDevice(id, "", &senderId)
		assert.Error(t, err)
	})
}
