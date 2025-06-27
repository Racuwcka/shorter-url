package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/Racuwcka/shorter-url/internal/config"
	"github.com/Racuwcka/shorter-url/internal/handler/shortener/add"
	"github.com/Racuwcka/shorter-url/internal/handler/shortener/original"
	"github.com/Racuwcka/shorter-url/internal/handler/shortener/redirect"
	"github.com/Racuwcka/shorter-url/internal/handler/shortener/short"
	addService "github.com/Racuwcka/shorter-url/internal/service/add"
	originalService "github.com/Racuwcka/shorter-url/internal/service/original"
	redirectService "github.com/Racuwcka/shorter-url/internal/service/redirect"
	shortService "github.com/Racuwcka/shorter-url/internal/service/short"
	"github.com/Racuwcka/shorter-url/internal/storage/cache"
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

	http.Handle("/swagger/", http.StripPrefix("/swagger/", http.FileServer(http.Dir("./../../swagger-ui"))))

	http.HandleFunc("/swagger", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/swagger/swagger.html", http.StatusFound)
	})

	repo := cache.NewShortenCache(cfg.Capacity)

	addHandler := add.New(addService.New(baseUrl, repo))
	http.HandleFunc("POST /api/v1/shorten", addHandler.Handle)

	getShortHandler := short.New(shortService.NewGetShortService(baseUrl, repo))
	http.HandleFunc("GET /api/v1/shorten", getShortHandler.Handle)

	getOriginHandler := original.New(originalService.NewGetOriginalService(baseUrl, repo))
	http.HandleFunc("GET /api/v1/original", getOriginHandler.Handle)

	getHandler := redirect.New(redirectService.NewGetService(repo))
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
