package app

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Racuwcka/shorter-url/internal/config"
	"github.com/Racuwcka/shorter-url/internal/router"
	"github.com/Racuwcka/shorter-url/pkg/closer"
)

func Run(ctx context.Context) error {
	cfg := config.MustLoadConfig()
	shutdowner := &closer.Closer{}

	addr := ":" + cfg.Port
	srv := &http.Server{
		Addr:    addr,
		Handler: router.New(cfg),
	}

	shutdowner.Add(srv.Shutdown)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen and serve: %v", err)
		}
	}()

	<-ctx.Done()
	log.Println("shutting down server gracefully")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout*time.Second)
	defer cancel()

	return shutdowner.Close(shutdownCtx)
}
