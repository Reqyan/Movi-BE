package main

import (
	"Movi-BE/controllers"
	"Movi-BE/models"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	models.InitDatabase()

	handler := controllers.New()

	server := &http.Server{
		Addr:    "0.0.0.0:8008",
		Handler: handler,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("ListenAndServe() failed: %v", err)
	}
}
