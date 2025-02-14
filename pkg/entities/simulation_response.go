package entities

type SimulationResponse struct {
	TotalAmount         float64 `json:"total_amount"`
	MonthlyInstallments float64 `json:"monthly_installments"`
	TotalInterest       float64 `json:"total_interest"`
}
