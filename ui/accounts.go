package ui

import (
	"github.com/gin-gonic/gin"
	"github.com/opencrypter/api/application"
	"github.com/opencrypter/api/infrastructure"
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

	var dto AccountDto
	err := context.BindJSON(&dto)
	if err == nil {
		err = service.Execute(dto.Id, dto.ExchangeId, dto.Name, dto.ApiKey, dto.ApiSecret)
	}
	if err != nil {
		apiError(context, err)
		return
	}
	apiSuccess(context, http.StatusOK, "")
}

func GetAccount(context *gin.Context) {
	context.JSON(http.StatusNotImplemented, Error{Message: "Not implemented"})
}

func GetBalances(context *gin.Context) {
	context.JSON(http.StatusNotImplemented, Error{Message: "Not implemented"})
}
