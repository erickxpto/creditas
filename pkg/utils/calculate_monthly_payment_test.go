package utils_test

import (
	"creditas/pkg/utils"
	"testing"
)

func TestCalculateMonthlyPayment(t *testing.T) {
	tests := []struct {
		pv         float64
		annualRate float64
		months     int
		expected   float64
	}{
		{1000, 5, 12, 85.6},
		{2000, 3, 24, 85.96},
		{1500, 2, 36, 42.96},
	}

	for _, test := range tests {
		result := utils.CalculateMonthlyPayment(test.pv, test.annualRate, test.months)
		if utils.RoundDecimal(result) != test.expected {
			t.Errorf("expected %v, got %v", test.expected, utils.RoundDecimal(result))
		}
	}
}
