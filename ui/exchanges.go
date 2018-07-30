package ui

import (
	"github.com/gin-gonic/gin"
	"github.com/opencrypter/api/application"
	"github.com/opencrypter/api/infrastructure"
	"net/http"
)

func GetExchanges(context *gin.Context) {
	service := application.NewGetAllExchanges(infrastructure.NewExchangeRepository())
	context.JSON(http.StatusOK, service.Execute())
}

func GetExchangeDetail(context *gin.Context) {
	context.JSON(http.StatusNotImplemented, Error{Message: "Not implemented"})
}
