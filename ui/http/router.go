package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/opencrypter/core/infrastructure"
	"net/http"
	"strings"
	"time"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors())

	router.POST("/devices", PostDevice)

	router.Group("/devices/:id", authMiddleware()).
		PATCH("", UpdateSenderId)

	router.Group("/accounts", authMiddleware()).
		GET("", GetAllAccounts).
		PUT("/:id", PutAccount).
		GET("/:id", GetAccount).
		GET("/:id/balances", GetBalances)

	router.Group("/exchanges", authMiddleware()).
		GET("", GetExchanges).
		GET("/:id", GetExchangeDetail).
		GET("/:id/tickers", GetExchangeTickers)

	router.Group("/tickers").
		GET("/:id/alerts", GetTickerAlerts).
		PUT("/:id/alerts/:alertId", PutTickerAlert).
		DELETE("/:id/alerts/:alertId", DeleteTickerAlert)

	return router
}

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Next()
	}
}

func authMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		request := context.Request
		validator := NewSignValidator(infrastructure.NewDeviceRepository())

		id := request.Header.Get("X-Api-Id")
		signature := request.Header.Get("X-Signature")
		date := request.Header.Get("Date")

		if id == "" || signature == "" || date == "" {
			context.AbortWithError(http.StatusForbidden, errors.New("authentication required"))
			return
		}

		dateTime, _ := time.Parse(time.RFC1123, date)
		body := string(readBody(context))
		payload := request.Method + request.URL.Path + request.URL.RawQuery + body + date
		payload = strings.Replace(payload, "\n", "", -1)
		payload = strings.Replace(payload, "\t", "", -1)
		payload = strings.Replace(payload, " ", "", -1)

		if err := validator.Validate(id, dateTime, payload, signature); err != nil {
			context.AbortWithError(http.StatusForbidden, err)
		}
	}
}
