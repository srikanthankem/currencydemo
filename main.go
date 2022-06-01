package main

import (
	"context"
	"currency-app/controllers"
	"currency-app/services"
	"log"

	"github.com/gin-gonic/gin"
)

var (
	server             *gin.Engine
	currencyService    services.CurrencyService
	currencyController controllers.CurrencyController
	ctx                context.Context
	err                error
)

func init() {
	ctx = context.TODO()
	currencyService = services.New(ctx)
	currencyController = controllers.New(currencyService)
	server = gin.Default()
}

func main() {
	basepath := server.Group("/v1")
	currencyController.RegisterCurrencyRoutes(basepath)
	log.Fatal(server.Run(":9090"))
}
