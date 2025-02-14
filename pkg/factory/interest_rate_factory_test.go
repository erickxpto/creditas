package factory

import (
	"creditas/pkg/services"
	"testing"
)

func TestCreateStrategy(t *testing.T) {
	factory := &InterestRateFactory{}

	tests := []struct {
		rateType string
		expected services.InterestRateStrategy
	}{
		{"variable", &services.VariableInterestRateStrategy{}},
		{"default", &services.DefaultInterestRateStrategy{}},
		{"unknown", &services.DefaultInterestRateStrategy{}}, // Default case
	}

	for _, test := range tests {
		result := factory.CreateStrategy(test.rateType)
		if _, ok := result.(*services.VariableInterestRateStrategy); ok && test.rateType == "variable" {
			continue
		}
		if _, ok := result.(*services.DefaultInterestRateStrategy); ok && test.rateType != "variable" {
			continue
		}
		t.Errorf("expected %T, got %T for rateType %s", test.expected, result, test.rateType)
	}
}
