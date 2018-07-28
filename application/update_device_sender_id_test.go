package application_test

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/opencrypter/api/application"
	"github.com/opencrypter/api/domain"
	"github.com/opencrypter/api/mock"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUpdateDeviceSenderId_Execute(t *testing.T) {
	mockedRepository := mock.NewMockDeviceRepository(gomock.NewController(t))
	service := application.NewUpdateDeviceSenderId(mockedRepository)

	newSenderId := "new"
	senderId := "old"
	os := "ios"
	secret := "test"
	device := &domain.Device{ID: uuid.NewV4().String(), Os: &os, SenderId: &senderId, Secret: &secret}

	t.Run("It should update the sender id", func(t *testing.T) {
		expected := &domain.Device{ID: device.ID, Os: device.Os, SenderId: &newSenderId, Secret: device.Secret}

		mockedRepository.
			EXPECT().
			DeviceOfId(device.ID).
			Return(device, nil)

		mockedRepository.
			EXPECT().
			Update(expected)

		err := service.Execute(device.ID, newSenderId)
		assert.NoError(t, err)
	})

	t.Run("It should return an error on missing device", func(t *testing.T) {
		id := uuid.NewV4().String()
		mockedRepository.
			EXPECT().
			DeviceOfId(id).
			Return(nil, errors.New("error"))

		err := service.Execute(id, "senderId")
		assert.Error(t, err)
	})

	t.Run("It should return an error on repository fail", func(t *testing.T) {
		expected := &domain.Device{ID: device.ID, Os: device.Os, SenderId: &newSenderId, Secret: device.Secret}

		mockedRepository.
			EXPECT().
			DeviceOfId(device.ID).
			Return(device, nil)

		mockedRepository.
			EXPECT().
			Update(expected).
			Return(errors.New("error"))

		err := service.Execute(device.ID, newSenderId)
		assert.Error(t, err)
	})
}
