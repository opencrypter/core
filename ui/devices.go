package ui

import (
	"github.com/gin-gonic/gin"
	"github.com/opencrypter/core/application"
	"github.com/opencrypter/core/domain"
	"github.com/opencrypter/core/infrastructure"
	"net/http"
)

type DeviceDto struct {
	Id       string `binding:"required"`
	Os       string `binding:"required"`
	SenderId string
}

type SenderIdDto struct {
	SenderId string `binding:"required"`
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

func UpdateSenderId(context *gin.Context) {
	service := application.NewUpdateDeviceSenderId(infrastructure.NewDeviceRepository())
	var dto SenderIdDto

	err := context.BindJSON(&dto)
	if err == nil {
		err = service.Execute(context.Param("id"), dto.SenderId)
	}
	if err != nil {
		apiError(context, err)
		return
	}
	apiSuccess(context, http.StatusOK, "")
}
