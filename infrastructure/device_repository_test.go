package infrastructure

import (
	"github.com/opencrypter/api/domain"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGormDeviceRepository_Save(t *testing.T) {
	repository := NewDeviceRepository()
	os := "ios"
	senderId := "abc"

	t.Run("It should save a new device", func(t *testing.T) {
		id, _ := uuid.NewV4()
		device, _ := domain.NewDevice(id.String(), os, &senderId)
		repository.Add(device)
		savedDevice, _ := repository.DeviceOfId(device.ID)
		assert.NotNil(t, savedDevice)
	})

	t.Run("It should save a device without sender id", func(t *testing.T) {
		id, _ := uuid.NewV4()
		device, _ := domain.NewDevice(id.String(), os, nil)
		repository.Add(device)
		savedDevice, _ := repository.DeviceOfId(device.ID)
		assert.NotNil(t, savedDevice)
	})

	t.Run("It should return error on duplicate", func(t *testing.T) {
		id, _ := uuid.NewV4()
		device, _ := domain.NewDevice(id.String(), os, &senderId)
		repository.Add(device)
		err := repository.Add(device)
		assert.Error(t, err)
	})

	t.Run("It should return error on invalid device", func(t *testing.T) {
		device := &domain.Device{}
		repository.Add(device)
		err := repository.Add(device)
		assert.Error(t, err)
	})
}
