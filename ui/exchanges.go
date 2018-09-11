package ui

import (
	"github.com/gin-gonic/gin"
	"github.com/opencrypter/core/application"
	"github.com/opencrypter/core/infrastructure"
	"net/http"
)

// Get exchanges.
// Returns all available exchanges.
func GetExchanges(context *gin.Context) {
	service := application.NewGetAllExchanges(infrastructure.NewExchangeRepository())
	context.JSON(http.StatusOK, service.Execute())
}

// Get exchange detail.
// Returns the detail of an exchange.
func GetExchangeDetail(context *gin.Context) {
	context.JSON(http.StatusNotImplemented, Error{Message: "Not implemented"})
}
