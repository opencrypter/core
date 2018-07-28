package infrastructure_test

import (
	"github.com/opencrypter/api/domain"
	"github.com/opencrypter/api/infrastructure"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGormDeviceRepository_Save(t *testing.T) {
	repository := infrastructure.NewDeviceRepository()
	os := "ios"
	senderId := "abc"

	t.Run("It should save a new device", func(t *testing.T) {
		device, _ := domain.NewDevice(uuid.NewV4().String(), os, &senderId)
		repository.Add(device)
		savedDevice, _ := repository.DeviceOfId(device.ID)
		assert.NotNil(t, savedDevice)
	})

	t.Run("It should save a device without sender id", func(t *testing.T) {
		device, _ := domain.NewDevice(uuid.NewV4().String(), os, nil)
		repository.Add(device)
		savedDevice, _ := repository.DeviceOfId(device.ID)
		assert.NotNil(t, savedDevice)
	})

	t.Run("It should return error on duplicate", func(t *testing.T) {
		device, _ := domain.NewDevice(uuid.NewV4().String(), os, &senderId)
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

func TestGormDeviceRepository_DeviceOfId(t *testing.T) {
	repository := infrastructure.NewDeviceRepository()
	t.Run("It should find device", func(t *testing.T) {
		senderId := "abc"
		os := "ios"
		secret := "test"
		existingDevice := &domain.Device{ID: uuid.NewV4().String(), Secret: &secret, Os: &os, SenderId: &senderId}
		repository.Add(existingDevice)

		device, _ := repository.DeviceOfId(existingDevice.ID)
		assert.NotNil(t, device)
	})

	t.Run("It should return error on missing device", func(t *testing.T) {
		_, err := repository.DeviceOfId(uuid.NewV4().String())
		assert.Error(t, err)
	})
}

func TestGormDeviceRepository_Update(t *testing.T) {
	repository := infrastructure.NewDeviceRepository()
	t.Run("It should update device", func(t *testing.T) {
		senderId := "abc"
		os := "ios"
		secret := "test"
		existingDevice := &domain.Device{ID: uuid.NewV4().String(), Secret: &secret, Os: &os, SenderId: &senderId}
		repository.Add(existingDevice)

		otherSenderId := "other-sender-id"
		existingDevice.SenderId = &otherSenderId
		err := repository.Update(existingDevice)

		updatedDevice, _ := repository.DeviceOfId(existingDevice.ID)
		assert.NoError(t, err)
		assert.Equal(t, &otherSenderId, updatedDevice.SenderId)
	})

	t.Run("It should return error on invalid device", func(t *testing.T) {
		os := "ios"
		device := &domain.Device{ID: uuid.NewV4().String(), Os: &os}

		err := repository.Update(device)
		assert.Error(t, err)
	})
}
