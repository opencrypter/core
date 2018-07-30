package ui_test

import (
	"bytes"
	"encoding/json"
	"github.com/opencrypter/api/domain"
	"github.com/opencrypter/api/infrastructure"
	"github.com/opencrypter/api/ui"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPostDevice(t *testing.T) {
	router := ui.NewRouter()
	senderId := "abc"
	os := "ios"
	t.Run("It should create a device", func(t *testing.T) {
		device := domain.Device{
			ID:       uuid.NewV4().String(),
			Os:       &os,
			SenderId: &senderId,
		}

		recorder := httptest.NewRecorder()
		buffer := new(bytes.Buffer)
		json.NewEncoder(buffer).Encode(&device)
		request, _ := http.NewRequest("POST", "/devices", buffer)
		request.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusCreated, recorder.Code)
	})

	t.Run("It should return bad request on invalid device", func(t *testing.T) {
		device := domain.Device{
			ID:       "invalid",
			Os:       &os,
			SenderId: &senderId,
		}

		recorder := httptest.NewRecorder()
		buffer := new(bytes.Buffer)
		json.NewEncoder(buffer).Encode(&device)
		request, _ := http.NewRequest("POST", "/devices", buffer)
		request.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("It should return bad request on missing body", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		request, _ := http.NewRequest("POST", "/devices", new(bytes.Buffer))
		request.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("It should return conflict status code on duplicated device", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		buffer := new(bytes.Buffer)
		json.NewEncoder(buffer).Encode(suite.existingDevice)
		duplicatedRequest, _ := http.NewRequest("POST", "/devices", buffer)
		duplicatedRequest.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(recorder, duplicatedRequest)

		assert.Equal(t, http.StatusConflict, recorder.Code)
	})
}

func TestUpdateSenderId(t *testing.T) {
	router := ui.NewRouter()
	dto := &ui.SenderIdDto{SenderId: "sender-id"}
	t.Run("It should update the sender id", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		buffer := new(bytes.Buffer)
		json.NewEncoder(buffer).Encode(dto)
		request := suite.newAuthenticatedRequest(requestData{
			buffer: buffer,
			method: "PATCH",
			path:   "/devices/" + suite.existingDevice.ID,
		})
		router.ServeHTTP(recorder, request)

		device, _ := infrastructure.NewDeviceRepository().DeviceOfId(suite.existingDevice.ID)
		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Equal(t, &dto.SenderId, device.SenderId)
	})

	t.Run("It should return error on missing device", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		buffer := new(bytes.Buffer)
		json.NewEncoder(buffer).Encode(&dto)
		request := suite.newAuthenticatedRequest(requestData{
			buffer: buffer,
			method: "PATCH",
			path:   "/devices/" + uuid.NewV4().String(),
		})
		router.ServeHTTP(recorder, request)
		assert.Equal(t, http.StatusNotFound, recorder.Code)
	})

	t.Run("It should return a forbidden error on missing credentials", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		buffer := new(bytes.Buffer)
		json.NewEncoder(buffer).Encode(dto)
		request, _ := http.NewRequest("PATCH", "/devices/"+suite.existingDevice.ID, buffer)
		router.ServeHTTP(recorder, request)
		assert.Equal(t, http.StatusForbidden, recorder.Code)
	})

	t.Run("It should return a forbidden error on invalid signature", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		buffer := new(bytes.Buffer)
		json.NewEncoder(buffer).Encode(dto)
		request := suite.newAuthenticatedRequest(requestData{
			buffer: buffer,
			method: "PATCH",
			path:   "/devices/" + suite.existingDevice.ID,
		})
		request.Header.Set("X-Signature", "invalid")
		router.ServeHTTP(recorder, request)
		assert.Equal(t, http.StatusForbidden, recorder.Code)
	})
}
