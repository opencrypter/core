package application

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/opencrypter/api/domain"
	"github.com/opencrypter/api/mock"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateDevice_Execute(t *testing.T) {
	mockedRepository := mock.NewMockDeviceRepository(gomock.NewController(t))
	DeviceRepository = mockedRepository
	id, _ := uuid.NewV4()
	os := "ios"
	senderId := "abc"

	domain.GenerateSecret = func() string {
		return "test"
	}

	t.Run("It should create a new device", func(t *testing.T) {
		expectedDevice := &domain.Device{
			ID:       id.String(),
			Os:       os,
			SenderId: &senderId,
			Secret:   "test",
		}

		mockedRepository.
			EXPECT().
			Add(expectedDevice)

		device, _ := CreateDevice(id.String(), os, &senderId)
		assert.Equal(t, expectedDevice, device)
	})

	t.Run("It should return an error on repository fail", func(t *testing.T) {
		expectedDevice := &domain.Device{
			ID:       id.String(),
			Os:       os,
			SenderId: &senderId,
			Secret:   "test",
		}

		mockedRepository.
			EXPECT().
			Add(expectedDevice).
			Return(errors.New("error"))

		_, err := CreateDevice(id.String(), os, &senderId)
		assert.Error(t, err)
	})
}