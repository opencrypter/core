package ui

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func PostDevice(context *gin.Context) {
	context.JSON(http.StatusNotImplemented, Error{Message: "Not implemented"})
}
