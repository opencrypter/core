package domain

import (
	"time"
)

const maxTimeToExpire = 120

type ValidateSignature struct {
	repository DeviceRepository
}

func NewValidateSignature(repository DeviceRepository) ValidateSignature {
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
		return NewInvalidSignatureDateError(deviceId, date)
	}
	if serverTime.Sub(timeInUTC).Seconds() > maxTimeToExpire {
		return NewExpiredRequestError(deviceId)
	}
	return nil
}
