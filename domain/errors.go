package domain

import (
	"errors"
	"fmt"
)

type InvalidDeviceError struct{ error }
type DuplicatedDeviceError struct{ error }

// Error about invalid device.
func NewInvalidDeviceError(message string) InvalidDeviceError {
	return InvalidDeviceError{error: errors.New(message)}
}

// Error about duplicated device.
func NewDuplicatedDeviceError(id string) DuplicatedDeviceError {
	return DuplicatedDeviceError{errors.New(fmt.Sprintf("Device %s is duplicated", id))}
}
