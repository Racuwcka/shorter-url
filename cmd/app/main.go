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
	baseUrl := "http://localhost:8080"
	addHandler := handlers.NewAddHandler(services.NewAddService(repo))
	http.HandleFunc("POST /api/v1/shorten", addHandler.Handle)

	getShortHandler := handlers.NewGetShortHandler(services.NewGetShortService(baseUrl, repo))
	http.HandleFunc("GET /api/v1/shorten", getShortHandler.Handle)

	getOriginHandler := handlers.NewGetOriginalHandler(services.NewGetOriginalService(baseUrl, repo))
	http.HandleFunc("GET /api/v1/original", getOriginHandler.Handle)

	getHandler := handlers.NewGetHandler(services.NewGetService(repo))
	http.HandleFunc("GET /link/{shortID}", getHandler.Handle)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
