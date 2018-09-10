package ui

import (
	"github.com/gin-gonic/gin"
	"github.com/opencrypter/core/application"
	"github.com/opencrypter/core/infrastructure"
	"net/http"
)

type AccountDto struct {
	Id         string  `binding:"required"`
	ExchangeId *string `binding:"required"`
	Name       *string `binding:"required"`
	ApiKey     *string `binding:"required"`
	ApiSecret  *string `binding:"required"`
}

func PutAccount(context *gin.Context) {
	service := application.NewSaveAccount(infrastructure.NewAccountRepository())
	deviceId := context.Request.Header.Get("X-Api-Id")
	var dto AccountDto
	err := context.BindJSON(&dto)
	if err == nil {
		err = service.Execute(dto.Id, &deviceId, dto.ExchangeId, dto.Name, dto.ApiKey, dto.ApiSecret)
	}
	if err != nil {
		apiError(context, err)
		return
	}
	apiSuccess(context, http.StatusOK, "")
}

func GetAllAccounts(context *gin.Context) {
	service := application.NewGetAllAccounts(infrastructure.NewAccountRepository())
	deviceId := context.Request.Header.Get("X-Api-Id")
	devices := service.Execute(deviceId)
	apiSuccess(context, http.StatusOK, devices)
}

func GetAccount(context *gin.Context) {
	context.JSON(http.StatusNotImplemented, Error{Message: "Not implemented"})
}

func GetBalances(context *gin.Context) {
	service := application.NewGetBalances(infrastructure.NewAccountRepository())
	id := context.Param("id")
	balances, err := service.Execute(id)
	if err != nil {
		apiError(context, err)
		return
	}

	apiSuccess(context, http.StatusOK, balances)
}
