package domain

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"github.com/satori/go.uuid"
	"math/rand"
)

var GenerateSecret = func() *string {
	data := make([]byte, 10)
	for i := range data {
		data[i] = byte(rand.Intn(256))
	}
	encoded := base64.StdEncoding.EncodeToString(sha256.New().Sum(data))
	return &encoded
}

type Device struct {
	ID       string  `gorm:"primary_key;type:uuid";json:"id"`
	Os       *string `gorm:"type:varchar;not null";json:"os"`
	SenderId *string `json:"senderId"`
	Secret   *string `gorm:"type:varchar;not null";json:"secret"`
}

type DeviceRepository interface {
	Add(device *Device) error
	DeviceOfId(id string) (*Device, error)
	Update(device *Device) error
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
		Os:       &os,
		SenderId: senderId,
		Secret:   GenerateSecret(),
	}, nil
}

func (d Device) Sign(payload string) string {
	sig := hmac.New(sha256.New, []byte(*d.Secret))
	sig.Write([]byte(payload))

	return base64.StdEncoding.EncodeToString(sig.Sum(nil))
}

func (d Device) ValidateSignature(payload string, signature string) error {
	if d.Sign(payload) != signature {
		return NewInvalidSignatureError(d.ID, signature)
	}
	return nil
}
