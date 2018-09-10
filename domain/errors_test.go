package domain_test

import (
	"fmt"
	"github.com/opencrypter/core/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewInvalidDeviceError(t *testing.T) {
	message := "error"
	assert.Equal(t, message, domain.NewInvalidDeviceError(message).Error())
}

func TestNewDuplicatedDeviceError(t *testing.T) {
	id := "cfe64055-5059-4464-8c02-6f5500cbaf08"
	assert.Regexp(t, fmt.Sprintf(".*(%s).*", id), domain.NewDuplicatedDeviceError(id).Error())
}

func TestNewAccountNotFoundError(t *testing.T) {
	id := "cfe64055-5059-4464-8c02-6f5500cbaf08"
	assert.Regexp(t, fmt.Sprintf(".*(%s).*", id), domain.NewAccountNotFoundError(id).Error())
}
