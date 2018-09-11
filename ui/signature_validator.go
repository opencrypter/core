package ui

import (
	"github.com/opencrypter/core/domain"
	"time"
)

const maxTimeToExpire = 120

type ValidateSignature struct {
	repository domain.DeviceRepository
}

func NewSignValidator(repository domain.DeviceRepository) ValidateSignature {
	return ValidateSignature{repository}
}

func (v ValidateSignature) Validate(deviceId string, date time.Time, payload string, signature string) error {
	err := v.validateDate(date, deviceId)
	if err != nil {
		return err
	}

	device, err := v.repository.DeviceOfId(deviceId)
	if err != nil {
		return err
	}

	return device.ValidateSignature(payload, signature)
}

func (v ValidateSignature) validateDate(date time.Time, deviceId string) error {
	timeInUTC := date.In(time.UTC)
	serverTime := time.Now().In(time.UTC)

	if timeInUTC.After(serverTime) {
		return domain.NewInvalidSignatureDateError(deviceId, date)
	}
	if serverTime.Sub(timeInUTC).Seconds() > maxTimeToExpire {
		return domain.NewExpiredRequestError(deviceId)
	}
	return nil
}
