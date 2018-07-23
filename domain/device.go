package domain

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/satori/go.uuid"
	"math/rand"
)

var GenerateSecret = func() string {
	data := make([]byte, 10)
	for i := range data {
		data[i] = byte(rand.Intn(256))
	}
	return hex.EncodeToString(sha256.New().Sum(data))
}

type Device struct {
	ID       string  `json:"id",gorm:"primary_key;type:uuid"`
	Os       string  `json:"os",gorm:"type:varchar;not null"`
	SenderId *string `json:"senderId"`
	Secret   string  `json:"secret",gorm:"type:varchar;not null"`
}

type DeviceRepository interface {
	Add(device *Device) error
	DeviceOfId(id string) (*Device, error)
}

func NewDevice(id string, os string, senderId *string) (*Device, error) {
	if _, err := uuid.FromString(id); err != nil {
		return nil, NewInvalidDeviceError(err.Error())
	}
	if os == "" {
		return nil, NewInvalidDeviceError("Device OS cannot be empty")
	}
	return &Device{
		ID:       id,
		Os:       os,
		SenderId: senderId,
		Secret:   GenerateSecret(),
	}, nil
}
