package handlers

import (
	"creditas/pkg/entities"
	"creditas/pkg/factory"
	"creditas/pkg/services"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

// SimulateHandler handles simulation requests
type SimulateHandler struct {
	factory           *factory.InterestRateFactory
	currencyConverter services.CurrencyConverter
	emailService      *services.EmailService
}

// NewSimulateHandler creates a new SimulateHandler
func NewSimulateHandler(emailService *services.EmailService) *SimulateHandler {
	return &SimulateHandler{
		factory:           &factory.InterestRateFactory{},
		currencyConverter: &services.DefaultCurrencyConverter{},
		emailService:      emailService,
	}
}

// ServeHTTP handles HTTP requests for simulation
// @Summary Simulate loan
// @Description Simulate a loan with given parameters
// @Tags simulate
// @Accept  json
// @Produce  json
// @Param   request body entities.SimulationRequest true "Simulation Request"
// @Success 200 {array} entities.SimulationResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /simulate [post]
func (h *SimulateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	var reqs []entities.SimulationRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&reqs); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	responses := make([]entities.SimulationResponse, len(reqs))
	var wg sync.WaitGroup
	for i, req := range reqs {
		wg.Add(1)
		go func(i int, req entities.SimulationRequest) {
			defer wg.Done()
			strategy := h.factory.CreateStrategy(req.InterestRateType)
			if strategy == nil {
				http.Error(w, "Invalid interest rate type", http.StatusBadRequest)
				return
			}
			service := services.NewSimulationService(strategy, h.currencyConverter)
			responses[i] = service.Simulate(req)

			// Enviar e-mail após a simulação
			// A verificação de erro somente faz o log que o envio falhou para que não
			// tenhamos problemas ao travar a aplicação com relação aos testes.
			// O ideal é que esse erro fosse tratado, talvez enviado para uma fila para que pudesse reprocessar o envio.
			subject := "Loan Simulation Result"
			body := fmt.Sprintf("Total Amount: %.2f\nMonthly Installments: %.2f\nTotal Interest: %.2f",
				responses[i].TotalAmount, responses[i].MonthlyInstallments, responses[i].TotalInterest)
			err := h.emailService.SendEmail(req.Email, subject, body)
			if err != nil {
				log.Println("Failed to send email")
			}

		}(i, req)
	}
	wg.Wait()

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(responses)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
