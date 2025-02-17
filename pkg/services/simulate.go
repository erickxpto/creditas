package services

import (
	"creditas/pkg/entities"
	"creditas/pkg/utils"
	"time"
)

// InterestRateStrategy define a interface para estratégias de taxa de juros
type InterestRateStrategy interface {
	GetAnnualRate(age int) float64
}

// DefaultInterestRateStrategy implementa uma estratégia padrão para taxas de juros
type DefaultInterestRateStrategy struct{}

// GetAnnualRate retorna a taxa de juros anual baseada na idade
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

// VariableInterestRateStrategy implementa uma estratégia variável para taxas de juros
type VariableInterestRateStrategy struct{}

// GetAnnualRate retorna a taxa de juros anual baseada na idade com um componente variável
func (s *VariableInterestRateStrategy) GetAnnualRate(age int) float64 {
	// Aqui, poderia ser chamado um serviço para obter a taxa variável
	baseRate := 3.5
	return baseRate + (float64(age) / 100)
}

// SimulationService lida com simulações de empréstimo
type SimulationService struct {
	rateStrategy      InterestRateStrategy
	currencyConverter CurrencyConverter
}

// NewSimulationService cria um novo SimulationService
func NewSimulationService(strategy InterestRateStrategy, converter CurrencyConverter) *SimulationService {
	return &SimulationService{
		rateStrategy:      strategy,
		currencyConverter: converter,
	}
}

// Simulate realiza a simulação de empréstimo
func (s *SimulationService) Simulate(req entities.SimulationRequest) entities.SimulationResponse {
	birthday, err := parseDate(req.Birthday)
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

// parseDate analisa uma string de data no formato "2006-01-02"
func parseDate(dateStr string) (time.Time, error) {
	const layout = "2006-01-02"
	return time.Parse(layout, dateStr)
}
