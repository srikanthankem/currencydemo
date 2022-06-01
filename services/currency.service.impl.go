package services

import (
	"context"
	"currency-app/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var client *http.Client

type CurrencyServiceImpl struct {
	ctx context.Context
}

func New(ctx context.Context) CurrencyService {
	return &CurrencyServiceImpl{
		ctx: ctx,
	}
}

func (c *CurrencyServiceImpl) GetCurrency(name *string) (*models.Currency, error) {
	var currency *models.Currency
	response, err := http.Get("https://api.hitbtc.com/api/3/public/currency/" + *name)

	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &currency); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}
	return currency, err
}

func (u *CurrencyServiceImpl) GetAll() (*models.Currencies, error) {
	var currencies *models.Currencies
	response, err := http.Get("https://api.hitbtc.com/api/3/public/currency/")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &currencies); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	return currencies, err

}
