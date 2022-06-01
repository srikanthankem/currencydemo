package services

import "currency-app/models"

type CurrencyService interface {
	GetCurrency(*string) (*models.Currency, error)
	GetAll() (*models.Currencies, error)
}
