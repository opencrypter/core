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

func TestCreateDevice_Execute(t *testing.T) {
	mockedRepository := mock.NewMockDeviceRepository(gomock.NewController(t))
	application.DeviceRepository = mockedRepository
	id := uuid.NewV4().String()
	os := "ios"
	senderId := "abc"
	secret := "test"
	domain.GenerateSecret = func() *string {
		return &secret
	}

	t.Run("It should create a new device", func(t *testing.T) {
		expectedDevice := &domain.Device{
			ID:       id,
			Os:       &os,
			SenderId: &senderId,
			Secret:   &secret,
		}

		mockedRepository.
			EXPECT().
			Add(expectedDevice)

		device, _ := application.CreateDevice(id, os, &senderId)
		assert.Equal(t, expectedDevice, device)
	})

	t.Run("It should return an error on repository fail", func(t *testing.T) {
		expectedDevice := &domain.Device{
			ID:       id,
			Os:       &os,
			SenderId: &senderId,
			Secret:   &secret,
		}

		mockedRepository.
			EXPECT().
			Add(expectedDevice).
			Return(errors.New("error"))

		_, err := application.CreateDevice(id, os, &senderId)
		assert.Error(t, err)
	})
}
