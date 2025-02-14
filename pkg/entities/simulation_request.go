package entities

type SimulationRequest struct {
	Amount           float64 `json:"amount"`
	Birthday         string  `json:"birthday"`
	PaymentTerm      int     `json:"payment_term"`
	Email            string  `json:"email"`
	InterestRateType string  `json:"interest_rate_type"`
	Currency         string  `json:"currency"`
}
