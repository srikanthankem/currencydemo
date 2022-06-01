package controllers

import (
	"currency-app/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CurrencyController struct {
	CurrencyService services.CurrencyService
}

func New(currencyService services.CurrencyService) CurrencyController {
	return CurrencyController{
		CurrencyService: currencyService,
	}
}

func (cc *CurrencyController) GetCurrency(ctx *gin.Context) {
	currencyName := ctx.Param("name")
	currency, err := cc.CurrencyService.GetCurrency(&currencyName)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, currency)
}

func (cc *CurrencyController) GetAll(ctx *gin.Context) {
	currencies, err := cc.CurrencyService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, currencies)
}

func (cc *CurrencyController) RegisterCurrencyRoutes(rg *gin.RouterGroup) {
	currencyroute := rg.Group("/currency")
	currencyroute.GET("/get/:name", cc.GetCurrency)
	currencyroute.GET("/getall", cc.GetAll)
}
