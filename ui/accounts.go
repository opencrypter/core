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

// Puts an account.
// If there is an existing resource with the same ID, it will be overwritten with the new values.
func PutAccount(context *gin.Context) {
	service := application.NewSaveAccount(infrastructure.NewAccountRepository())
	deviceId := context.Request.Header.Get("X-Api-Id")
	var dto AccountDto
	if err := context.BindJSON(&dto); err != nil {
		apiError(context, err)
		return
	}
	if err := service.Execute(dto.Id, &deviceId, dto.ExchangeId, dto.Name, dto.ApiKey, dto.ApiSecret); err != nil {
		apiError(context, err)
		return
	}
	apiSuccess(context, http.StatusOK, "")
}

// Get all accounts.
// Returns all persisted accounts related with a device.
func GetAllAccounts(context *gin.Context) {
	service := application.NewGetAllAccounts(infrastructure.NewAccountRepository())
	deviceId := context.Request.Header.Get("X-Api-Id")
	devices := service.Execute(deviceId)
	apiSuccess(context, http.StatusOK, devices)
}

// Get account.
// Returns the detail of an account.
func GetAccount(context *gin.Context) {
	context.JSON(http.StatusNotImplemented, Error{Message: "Not implemented"})
}

// Get balances.
// Returns all balances related with a device.
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
