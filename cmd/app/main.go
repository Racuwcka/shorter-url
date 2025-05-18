package main

import (
	"github.com/Racuwcka/shorter-url/internal/pkg/handlers"
	"github.com/Racuwcka/shorter-url/internal/pkg/repositories"
	"github.com/Racuwcka/shorter-url/internal/pkg/services"
	"log"
	"net/http"
)

func main() {
	repo := repositories.NewShortenCache()
	handler := handlers.NewHandler(services.NewAddService(repo))
	http.HandleFunc("/api/v1/shorten", handler.Handle)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
