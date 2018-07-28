package ui_test

import (
	"bytes"
	"encoding/json"
	"github.com/opencrypter/api/domain"
	"github.com/opencrypter/api/infrastructure"
	"github.com/opencrypter/api/ui"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestPostDevice(t *testing.T) {
	router := ui.NewRouter()
	senderId := "abc"
	t.Run("It should create a device", func(t *testing.T) {
		device := domain.Device{
			ID:       uuid.NewV4().String(),
			Os:       "ios",
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
			Os:       "ios",
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
		device := domain.Device{
			ID:       uuid.NewV4().String(),
			Os:       "ios",
			SenderId: &senderId,
		}

		buffer := new(bytes.Buffer)
		json.NewEncoder(buffer).Encode(&device)
		request, _ := http.NewRequest("POST", "/devices", buffer)
		request.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(httptest.NewRecorder(), request)

		recorder := httptest.NewRecorder()
		buffer = new(bytes.Buffer)
		json.NewEncoder(buffer).Encode(&device)
		duplicatedRequest, _ := http.NewRequest("POST", "/devices", buffer)
		duplicatedRequest.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(recorder, duplicatedRequest)

		assert.Equal(t, http.StatusConflict, recorder.Code)
	})
}

func TestUpdateSenderId(t *testing.T) {
	router := ui.NewRouter()
	device := &domain.Device{ID: uuid.NewV4().String(), Os: "ios", Secret: "secret"}
	infrastructure.NewDeviceRepository().Add(device)

	t.Run("It should update the sender id", func(t *testing.T) {
		senderId := "sender-id"
		dto := ui.SenderIdDto{SenderId: senderId}

		recorder := httptest.NewRecorder()
		buffer := new(bytes.Buffer)
		json.NewEncoder(buffer).Encode(&dto)
		request, _ := http.NewRequest("PATCH", "/devices/"+device.ID, buffer)
		request.Header.Set("X-Api-Id", device.ID)
		request.Header.Set("Date", time.Now().Format(time.RFC1123))
		request.Header.Set("X-Signature", buildSignature(request, *device))
		router.ServeHTTP(recorder, request)

		device, _ := infrastructure.NewDeviceRepository().DeviceOfId(device.ID)

		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Equal(t, device.SenderId, &senderId)
	})

	t.Run("It should return error on missing device", func(t *testing.T) {
		dto := ui.SenderIdDto{SenderId: "sender-id"}

		recorder := httptest.NewRecorder()
		buffer := new(bytes.Buffer)
		json.NewEncoder(buffer).Encode(&dto)
		request, _ := http.NewRequest("PATCH", "/devices/"+uuid.NewV4().String(), buffer)
		request.Header.Set("X-Api-Id", device.ID)
		request.Header.Set("Date", time.Now().Format(time.RFC1123))
		request.Header.Set("X-Signature", buildSignature(request, *device))
		router.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusNotFound, recorder.Code)
	})
}

func buildSignature(request *http.Request, device domain.Device) string {
	bodyBytes, _ := ioutil.ReadAll(request.Body)
	request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	date := request.Header.Get("Date")
	payload := request.Method + request.URL.Path + request.URL.RawQuery + string(bodyBytes) + date
	payload = strings.Replace(payload, " ", "", -1)

	s := device.Sign(payload)
	return s
}
