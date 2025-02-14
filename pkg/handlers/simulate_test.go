package handlers_test

import (
	"bytes"
	"creditas/pkg/entities"
	"creditas/pkg/handlers"
	"creditas/pkg/services"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSimulateHandler(t *testing.T) {
	emailService := &services.EmailService{}
	handler := handlers.NewSimulateHandler(emailService)

	reqs := []entities.SimulationRequest{
		{
			Amount:           1000,
			Birthday:         "2000-01-01",
			PaymentTerm:      12,
			InterestRateType: "default",
			Currency:         "USD",
			Email:            "test@example.com",
		},
	}

	body, _ := json.Marshal(reqs)
	req, err := http.NewRequest("POST", "/simulate", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var responses []entities.SimulationResponse
	err = json.NewDecoder(rr.Body).Decode(&responses)
	if err != nil {
		t.Errorf("could not decode response: %v", err)
	}

	if len(responses) != 1 {
		t.Errorf("expected 1 response, got %v", len(responses))
		return
	}

	resp := responses[0]
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
