package services_test

import (
	"creditas/pkg/entities"
	"creditas/pkg/services"
	"testing"
)

func TestSimulate(t *testing.T) {
	strategy := &services.DefaultInterestRateStrategy{}
	converter := &services.DefaultCurrencyConverter{}
	service := services.NewSimulationService(strategy, converter)

	req := entities.SimulationRequest{
		Amount:           1000,
		Birthday:         "2000-01-01",
		PaymentTerm:      12,
		InterestRateType: "default",
		Currency:         "USD",
	}

	resp := service.Simulate(req)

	if resp.TotalAmount == 0 {
		t.Errorf("expected total amount to be greater than 0")
	}
	if resp.MonthlyInstallments == 0 {
		t.Errorf("expected monthly installments to be greater than 0")
	}
	if resp.TotalInterest == 0 {
		t.Errorf("expected total interest to be greater than 0")
	}
}
