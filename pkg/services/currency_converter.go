package services

import "fmt"

type CurrencyConverter interface {
	ConvertToBRL(amount float64, currency string) (float64, error)
}

type DefaultCurrencyConverter struct{}

// Aqui, poderiamos chamar um serviço externo e manter os valores em cache (ajudando na performance e não indo diversas vezes no serviço)
func (c *DefaultCurrencyConverter) ConvertToBRL(amount float64, currency string) (float64, error) {
	switch currency {
	case "USD":
		return amount * 5.0, nil
	case "EUR":
		return amount * 6.0, nil
	case "BRL":
		return amount, nil
	default:
		return 0, fmt.Errorf("unsupported currency: %s", currency)
	}
}
