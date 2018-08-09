package domain

import (
	"errors"
	"fmt"
	"time"
)

type DeviceNotFoundError struct{ error }
type AccountNotFoundError struct{ error }
type InvalidDeviceError struct{ error }
type DuplicatedDeviceError struct{ error }
type InvalidSignatureError struct{ error }
type ExpiredRequestError struct{ error }
type InvalidSignatureDateError struct{ error }

// Error about missing device.
func NewDeviceNotFoundError(id string) DeviceNotFoundError {
	return DeviceNotFoundError{error: errors.New(fmt.Sprintf("device %s not found", id))}
}

// Error about missing account.
func NewAccountNotFoundError(id string) AccountNotFoundError {
	return AccountNotFoundError{error: errors.New(fmt.Sprintf("account %s not found", id))}
}

// Error about invalid device.
func NewInvalidDeviceError(message string) InvalidDeviceError {
	return InvalidDeviceError{error: errors.New(message)}
}

// Error about duplicated device.
func NewDuplicatedDeviceError(id string) DuplicatedDeviceError {
	return DuplicatedDeviceError{errors.New(fmt.Sprintf("device %s is duplicated", id))}
}

// Error about invalid signature.
func NewInvalidSignatureError(id string, signature string) InvalidSignatureError {
	return InvalidSignatureError{errors.New(fmt.Sprintf("device %s received invalid signature: %s", id, signature))}
}

// Error about expired requests.
func NewExpiredRequestError(deviceId string) ExpiredRequestError {
	return ExpiredRequestError{errors.New("expired request for device " + deviceId)}
}

// Error about expired requests.
func NewInvalidSignatureDateError(deviceId string, date time.Time) InvalidSignatureDateError {
	message := fmt.Sprintf("device %s signed with invalid date %s", deviceId, date.Format(time.RFC1123))
	return InvalidSignatureDateError{errors.New(message)}
}
