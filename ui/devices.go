package ui

import (
	"github.com/gin-gonic/gin"
	"github.com/opencrypter/api/application"
	"github.com/opencrypter/api/domain"
	"github.com/opencrypter/api/infrastructure"
	"net/http"
)

type DeviceDto struct {
	Id       string `binding:"required"`
	Os       string `binding:"required"`
	SenderId string
}

func PostDevice(context *gin.Context) {
	application.DeviceRepository = infrastructure.NewDeviceRepository()
	var dto DeviceDto
	var device *domain.Device

	err := context.BindJSON(&dto)
	if err == nil {
		device, err = application.CreateDevice(dto.Id, dto.Os, &dto.SenderId)
	}
	if err != nil {
		apiError(context, err)
		return
	}
	apiSuccess(context, http.StatusCreated, device)
}
