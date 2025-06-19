package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/Racuwcka/shorter-url/internal/handler/shortener/original"
	"github.com/Racuwcka/shorter-url/internal/handler/shortener/short"
	add2 "github.com/Racuwcka/shorter-url/internal/services/add"
	original2 "github.com/Racuwcka/shorter-url/internal/services/original"
	redirect2 "github.com/Racuwcka/shorter-url/internal/services/redirect"
	short2 "github.com/Racuwcka/shorter-url/internal/services/short"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/Racuwcka/shorter-url/internal/config"
	"github.com/Racuwcka/shorter-url/internal/handler/shortener/add"
	"github.com/Racuwcka/shorter-url/internal/handler/shortener/redirect"
	"github.com/Racuwcka/shorter-url/internal/repositories"
	"github.com/Racuwcka/shorter-url/pkg/closer"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	if err := runServer(ctx); err != nil {
		log.Fatal(err)
	}
}

func runServer(ctx context.Context) error {
	cfg := config.LoadConfig()
	var (
		srv = &http.Server{
			Addr: cfg.Addr,
		}
		c = &closer.Closer{}
	)

	var baseUrl string

	switch cfg.Env {
	case envLocal:
		baseUrl = fmt.Sprintf("http://%s", cfg.Addr)
	case envDev:
		baseUrl = fmt.Sprintf("https://%s", cfg.Addr)
	case envProd:
		baseUrl = fmt.Sprintf("https://%s", cfg.Addr)
	}

	repo := repositories.NewShortenCache(cfg.Capacity)

	addHandler := add.New(add2.New(baseUrl, repo))
	http.HandleFunc("POST /api/v1/shorten", addHandler.Handle)

	getShortHandler := short.New(short2.NewGetShortService(baseUrl, repo))
	http.HandleFunc("GET /api/v1/shorten", getShortHandler.Handle)

	getOriginHandler := original.New(original2.NewGetOriginalService(baseUrl, repo))
	http.HandleFunc("GET /api/v1/original", getOriginHandler.Handle)

	getHandler := redirect.New(redirect2.NewGetService(repo))
	http.HandleFunc("GET /link/{shortID}", getHandler.Handle)

	c.Add(srv.Shutdown)

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen and serve: %v", err)
		}
	}()

	log.Printf("listening on %s", cfg.Addr)
	<-ctx.Done()

	log.Println("shutting down server gracefully")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout*time.Second)
	defer cancel()

	if err := c.Close(shutdownCtx); err != nil {
		return fmt.Errorf("closer: %v", err)
	}

	return nil
}
