package ui

import (
	"github.com/gin-gonic/gin"
	"github.com/opencrypter/api/domain"
	"net/http"
)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func apiError(context *gin.Context, error error) {
	abortWithError(context, error)
	body := Error{
		Code:    context.Writer.Status(),
		Message: error.Error(),
	}
	context.JSON(body.Code, body)
}

func apiSuccess(context *gin.Context, statusCode int, body interface{}) {
	context.JSON(statusCode, body)
}

func abortWithError(context *gin.Context, err error) {
	var code int
	switch err.(type) {
	case domain.InvalidDeviceError:
		code = http.StatusBadRequest
	case domain.DuplicatedDeviceError:
		code = http.StatusConflict
	}

	context.AbortWithError(code, err)
}
