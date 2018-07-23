package application

import (
	"github.com/opencrypter/api/domain"
)

var DeviceRepository domain.DeviceRepository

func CreateDevice(id string, os string, senderId *string) (*domain.Device, error) {
	var err error

	device, err := domain.NewDevice(id, os, senderId)

	if err == nil {
		err = DeviceRepository.Add(device)
	}

	if err != nil {
		return nil, err
	}

	return device, nil
}