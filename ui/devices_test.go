package ui

import (
	"bytes"
	"encoding/json"
	"github.com/opencrypter/api/domain"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPostDevice(t *testing.T) {
	router := NewRouter()
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
