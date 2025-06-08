package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/Racuwcka/shorter-url/cmd/closer"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	handlers2 "github.com/Racuwcka/shorter-url/pkg/handlers"
	"github.com/Racuwcka/shorter-url/pkg/repositories"
	services2 "github.com/Racuwcka/shorter-url/pkg/services"
)

const (
	listenAddr      = "localhost:8080"
	shutdownTimeout = 5 * time.Second
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	if err := runServer(ctx); err != nil {
		log.Fatal(err)
	}
}

func runServer(ctx context.Context) error {
	var (
		srv = &http.Server{
			Addr: listenAddr,
		}
		c = &closer.Closer{}
	)

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

	c.Add(srv.Shutdown)

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen and serve: %v", err)
		}
	}()

	log.Printf("listening on %s", listenAddr)
	<-ctx.Done()

	log.Println("shutting down server gracefully")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := c.Close(shutdownCtx); err != nil {
		return fmt.Errorf("closer: %v", err)
	}

	return nil
}
