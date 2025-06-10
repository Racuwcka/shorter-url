package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/Racuwcka/shorter-url/internal/closer"
	"github.com/Racuwcka/shorter-url/internal/config"
	handlers2 "github.com/Racuwcka/shorter-url/pkg/handlers"
	"github.com/Racuwcka/shorter-url/pkg/repositories"
	services2 "github.com/Racuwcka/shorter-url/pkg/services"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

type Config struct {
	Env             string
	Addr            string
	ShutdownTimeout time.Duration
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	if err := runServer(ctx); err != nil {
		log.Fatal(err)
	}
}

func loadConfig() *Config {
	envConfig := config.LoadEnvFile(".env")

	defaultEnv := envConfig.GetEnvString("env", "local")
	defaultAddr := envConfig.GetEnvString("addr", "localhost:8080")
	defaultShutdownTimeout := envConfig.GetEnvInt("shutdownTimeout", 5)

	env := flag.String("env", defaultEnv, "the app env")
	addr := flag.String("addr", defaultAddr, "the address to connect to")
	shutdownTimeout := flag.Int("shutdownTimeout", defaultShutdownTimeout, "shutdownTimeout time")

	flag.Parse()

	return &Config{
		Env:             *env,
		Addr:            *addr,
		ShutdownTimeout: time.Duration(*shutdownTimeout) * time.Second,
	}
}

func runServer(ctx context.Context) error {
	cfg := loadConfig()
	var (
		srv = &http.Server{
			Addr: cfg.Addr,
		}
		c = &closer.Closer{}
	)

	repo := repositories.NewShortenCache()

	var baseUrl string
	fmt.Println(cfg)

	switch cfg.Env {
	case envLocal:
		baseUrl = fmt.Sprintf("http://%s", cfg.Addr)
	case envDev:
		baseUrl = fmt.Sprintf("https://%s", cfg.Addr)
	case envProd:
		baseUrl = fmt.Sprintf("https://%s", cfg.Addr)
	}

	fmt.Println(baseUrl)

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
