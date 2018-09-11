package main

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/opencrypter/core/domain"
	"io/ioutil"
	"net/http"
)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func apiSuccess(context *gin.Context, statusCode int, body interface{}) {
	context.JSON(statusCode, body)
}

func apiError(context *gin.Context, error error) {
	abortWithError(context, error)
	body := Error{
		Code:    context.Writer.Status(),
		Message: error.Error(),
	}
	context.JSON(body.Code, body)
}

func abortWithError(context *gin.Context, err error) {
	var code int
	switch err.(type) {
	case domain.InvalidDeviceError:
		code = http.StatusBadRequest
	case domain.DuplicatedDeviceError:
		code = http.StatusConflict
	case domain.DeviceNotFoundError, domain.AccountNotFoundError:
		code = http.StatusNotFound
	}

	context.AbortWithError(code, err)
}

func readBody(context *gin.Context) []byte {
	body, _ := ioutil.ReadAll(context.Request.Body)
	context.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	return body
}
