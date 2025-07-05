package app

import (
	"context"
	"github.com/Racuwcka/shorter-url/internal/router"
	"log"
	"net/http"
	"time"

	"github.com/Racuwcka/shorter-url/internal/config"
	"github.com/Racuwcka/shorter-url/pkg/closer"
)

func Run(ctx context.Context) error {
	cfg := config.LoadConfig()
	c := &closer.Closer{}

	srv := &http.Server{
		Addr:    cfg.Addr,
		Handler: router.New(cfg),
	}

	c.Add(srv.Shutdown)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen and serve: %v", err)
		}
	}()

	log.Printf("listening on %s", cfg.Addr)

	<-ctx.Done()
	log.Println("shutting down server gracefully")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout*time.Second)
	defer cancel()

	return c.Close(shutdownCtx)
}
