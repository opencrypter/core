package main

import (
	"github.com/gin-gonic/gin"
	"github.com/opencrypter/api/ui"
)

func main() {
	NewRouter().Run()
}

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.Use(Cors())

	router.POST("/devices", ui.PostDevice)
	router.PUT("/accounts/:id", ui.PutAccount)
	router.GET("/accounts/:id", ui.GetAccount)
	router.GET("/accounts/:id/balances", ui.GetBalances)
	router.GET("/exchanges", ui.GetExchanges)
	router.GET("/exchanges/:id", ui.GetExchangeDetail)
	router.GET("/exchanges/:id/tickers", ui.GetExchangeTickers)
	router.GET("/tickers/:id/alerts", ui.GetTickerAlerts)
	router.PUT("/tickers/:id/alerts/:alertId", ui.PutTickerAlert)
	router.DELETE("/tickers/:id/alerts/:alertId", ui.DeleteTickerAlert)

	return router
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Next()
	}
}
