package ui

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetExchanges(context *gin.Context) {
	context.JSON(http.StatusNotImplemented, Error{Message: "Not implemented"})
}

func GetExchangeDetail(context *gin.Context) {
	context.JSON(http.StatusNotImplemented, Error{Message: "Not implemented"})
}
