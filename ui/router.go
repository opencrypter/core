package ui

import (
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors())

	router.POST("/devices", PostDevice)
	router.PUT("/accounts/:id", PutAccount)
	router.GET("/accounts/:id", GetAccount)
	router.GET("/accounts/:id/balances", GetBalances)
	router.GET("/exchanges", GetExchanges)
	router.GET("/exchanges/:id", GetExchangeDetail)
	router.GET("/exchanges/:id/tickers", GetExchangeTickers)
	router.GET("/tickers/:id/alerts", GetTickerAlerts)
	router.PUT("/tickers/:id/alerts/:alertId", PutTickerAlert)
	router.DELETE("/tickers/:id/alerts/:alertId", DeleteTickerAlert)

	return router
}

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Next()
	}
}
