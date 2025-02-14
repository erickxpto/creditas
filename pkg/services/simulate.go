package services

import (
	"creditas/pkg/entities"
	"creditas/pkg/utils"
	"time"
)

type InterestRateStrategy interface {
	GetAnnualRate(age int) float64
}

type DefaultInterestRateStrategy struct{}

func (s *DefaultInterestRateStrategy) GetAnnualRate(age int) float64 {
	switch {
	case age <= 25:
		return 5.0
	case age <= 40:
		return 3.0
	case age <= 60:
		return 2.0
	default:
		return 4.0
	}
}

type VariableInterestRateStrategy struct{}

func (s *VariableInterestRateStrategy) GetAnnualRate(age int) float64 {
	// Aqui, poderia ser chamado um serviço para obter a taxa variável
	indice := 3.5
	return indice + (float64(age) / 100)
}

type SimulationService struct {
	rateStrategy      InterestRateStrategy
	currencyConverter CurrencyConverter
}

func NewSimulationService(strategy InterestRateStrategy, converter CurrencyConverter) *SimulationService {
	return &SimulationService{
		rateStrategy:      strategy,
		currencyConverter: converter,
	}
}

func (s *SimulationService) Simulate(req entities.SimulationRequest) entities.SimulationResponse {
	layout := "2006-01-02"
	birthday, err := time.Parse(layout, req.Birthday)
	if err != nil {
		return entities.SimulationResponse{}
	}

	age := utils.CalculateAge(birthday)
	annualRate := s.rateStrategy.GetAnnualRate(age)

	amountInBRL, err := s.currencyConverter.ConvertToBRL(req.Amount, req.Currency)
	if err != nil {
		return entities.SimulationResponse{}
	}

	monthlyPayment := utils.CalculateMonthlyPayment(amountInBRL, annualRate, req.PaymentTerm)
	totalAmount := monthlyPayment * float64(req.PaymentTerm)
	totalInterest := totalAmount - amountInBRL

	return entities.SimulationResponse{
		TotalAmount:         utils.RoundDecimal(totalAmount),
		MonthlyInstallments: utils.RoundDecimal(monthlyPayment),
		TotalInterest:       utils.RoundDecimal(totalInterest),
	}
}
