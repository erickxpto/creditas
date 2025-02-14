package services_test

import (
	"creditas/pkg/services"
	"testing"
)

func TestConvertToBRL(t *testing.T) {
	converter := &services.DefaultCurrencyConverter{}

	tests := []struct {
		amount   float64
		currency string
		expected float64
	}{
		{100, "USD", 500},
		{100, "EUR", 600},
		{100, "BRL", 100},
	}

	for _, test := range tests {
		result, err := converter.ConvertToBRL(test.amount, test.currency)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if result != test.expected {
			t.Errorf("expected %v, got %v", test.expected, result)
		}
	}
}
