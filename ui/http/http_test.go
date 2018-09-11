package main_test

import (
	"bytes"
	"github.com/opencrypter/core/domain"
	"github.com/opencrypter/core/infrastructure"
	"github.com/satori/go.uuid"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var suite testSuite

type testSuite struct {
	existingDevice *domain.Device
}

type requestData struct {
	device *domain.Device
	method string
	path   string
	buffer *bytes.Buffer
}

func init() {
	os := "test"
	secret := "D9pokn8rL/g29+OwxEKY/BwUmvv0yJlvuSQnrkHkZJuTTKSVmRt4UrhV"
	senderId := "sender-id"

	suite = testSuite{existingDevice: &domain.Device{
		ID:       uuid.NewV4().String(),
		Os:       &os,
		SenderId: &senderId,
		Secret:   &secret,
	}}
	infrastructure.NewDeviceRepository().Add(suite.existingDevice)
}

func (testSuite) newAuthenticatedRequest(request requestData) *http.Request {
	if request.device == nil {
		request.device = suite.existingDevice
	}
	if request.buffer == nil {
		request.buffer = new(bytes.Buffer)
	}
	infrastructure.NewDeviceRepository().Update(request.device)
	r, _ := http.NewRequest(request.method, request.path, request.buffer)
	r.Header.Set("X-Api-Id", request.device.ID)
	r.Header.Set("Date", time.Now().Format(time.RFC1123))
	r.Header.Set("X-Signature", buildSignature(r, *request.device))

	return r
}

func buildSignature(request *http.Request, device domain.Device) string {
	bodyBytes, _ := ioutil.ReadAll(request.Body)
	request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	date := request.Header.Get("Date")
	payload := request.Method + request.URL.Path + request.URL.RawQuery + string(bodyBytes) + date
	payload = strings.Replace(payload, "\n", "", -1)
	payload = strings.Replace(payload, "\t", "", -1)
	payload = strings.Replace(payload, " ", "", -1)
	s := device.Sign(payload)

	return s
}
