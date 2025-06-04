package main

import (
	"log"
	"net/http"

	handlers2 "github.com/Racuwcka/shorter-url/pkg/handlers"
	"github.com/Racuwcka/shorter-url/pkg/repositories"
	services2 "github.com/Racuwcka/shorter-url/pkg/services"
)

func main() {
	repo := repositories.NewShortenCache()
	baseUrl := "http://localhost:8080"
	addHandler := handlers2.NewAddHandler(services2.NewAddService(repo))
	http.HandleFunc("POST /api/v1/shorten", addHandler.Handle)

	getShortHandler := handlers2.NewGetShortHandler(services2.NewGetShortService(baseUrl, repo))
	http.HandleFunc("GET /api/v1/shorten", getShortHandler.Handle)

	getOriginHandler := handlers2.NewGetOriginalHandler(services2.NewGetOriginalService(baseUrl, repo))
	http.HandleFunc("GET /api/v1/original", getOriginHandler.Handle)

	getHandler := handlers2.NewGetHandler(services2.NewGetService(repo))
	http.HandleFunc("GET /link/{shortID}", getHandler.Handle)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
