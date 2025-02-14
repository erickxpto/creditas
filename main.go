package main

import (
	_ "creditas/docs" // Import the generated docs
	"creditas/pkg/handlers"
	"creditas/pkg/services"
	"log"
	"net/http"
	"os"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Creditas API
// @version 1.0
// @description This is a sample server for Creditas.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	emailService := services.NewEmailService(
		os.Getenv("SMTP_HOST"),
		os.Getenv("SMTP_PORT"),
		os.Getenv("SMTP_USERNAME"),
		os.Getenv("SMTP_PASSWORD"),
	)

	mux := http.NewServeMux()
	simulateHandler := handlers.NewSimulateHandler(emailService)
	mux.Handle("/simulate", simulateHandler)

	// Swagger endpoint
	mux.Handle("/swagger/", httpSwagger.WrapHandler)

	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
