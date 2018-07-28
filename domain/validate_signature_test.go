package domain_test

import (
	"github.com/golang/mock/gomock"
	"github.com/opencrypter/api/domain"
	"github.com/opencrypter/api/mock"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestValidateSignature_Validate(t *testing.T) {
	mockedRepository := mock.NewMockDeviceRepository(gomock.NewController(t))
	service := domain.NewValidateSignature(mockedRepository)
	secret := "thisIsMySecret"
	device := &domain.Device{ID: uuid.NewV4().String(), Secret: &secret}

	t.Run("It should validate a valid signature", func(t *testing.T) {
		payload := "a random payload"
		signature := device.Sign(payload)

		mockedRepository.EXPECT().DeviceOfId(device.ID).Return(device, nil)

		assert.NoError(t, service.Validate(device.ID, time.Now(), payload, signature))
	})

	t.Run("It should validate a valid signature with other timezone", func(t *testing.T) {
		payload := "a random payload"
		signature := device.Sign(payload)

		mockedRepository.EXPECT().DeviceOfId(device.ID).Return(device, nil)

		location, _ := time.LoadLocation("Asia/Shanghai")
		date := time.Now().In(location)

		assert.NoError(t, service.Validate(device.ID, date, payload, signature))
	})

	t.Run("It should return an error on expired request", func(t *testing.T) {
		date := time.Now().Add(-121 * time.Duration(time.Second))
		err := service.Validate(uuid.NewV4().String(), date, "", "")
		assert.Error(t, err)
		assert.IsType(t, domain.ExpiredRequestError{}, err)
	})

	t.Run("It should return an error on receive a future date", func(t *testing.T) {
		date := time.Now().Add(100 * time.Duration(time.Second))
		err := service.Validate(uuid.NewV4().String(), date, "", "")
		assert.Error(t, err)
		assert.IsType(t, domain.InvalidSignatureDateError{}, err)
	})

	t.Run("It should return an error on missing device", func(t *testing.T) {
		payload := "a random payload"
		signature := device.Sign(payload)

		mockedRepository.EXPECT().DeviceOfId(device.ID).Return(nil, domain.NewDeviceNotFoundError(device.ID))

		err := service.Validate(device.ID, time.Now(), payload, signature)
		assert.Error(t, err)
		assert.IsType(t, domain.DeviceNotFoundError{}, err)
	})

	t.Run("It should return an error on invalid signature", func(t *testing.T) {
		mockedRepository.EXPECT().DeviceOfId(device.ID).Return(device, nil)

		err := service.Validate(device.ID, time.Now(), "a random payload", "invalid")
		assert.Error(t, err)
		assert.IsType(t, domain.InvalidSignatureError{}, err)
	})
}
