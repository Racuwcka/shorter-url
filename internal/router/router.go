package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Racuwcka/shorter-url/internal/config"
	addHandler "github.com/Racuwcka/shorter-url/internal/handler/shortener/add"
	originalHandler "github.com/Racuwcka/shorter-url/internal/handler/shortener/original"
	redirectHandler "github.com/Racuwcka/shorter-url/internal/handler/shortener/redirect"
	shortHandler "github.com/Racuwcka/shorter-url/internal/handler/shortener/short"
	addService "github.com/Racuwcka/shorter-url/internal/service/add"
	shortService "github.com/Racuwcka/shorter-url/internal/service/short"
	"github.com/Racuwcka/shorter-url/internal/service/shortener"
	"github.com/Racuwcka/shorter-url/internal/storage"
	"github.com/Racuwcka/shorter-url/internal/storage/cache"
	"github.com/Racuwcka/shorter-url/internal/storage/db"
	"github.com/Racuwcka/shorter-url/pkg/client/postgresql"
)

func New(cfg *config.Config) http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/swagger/", http.StripPrefix("/swagger/", http.FileServer(http.Dir("./../../swagger-ui"))))
	mux.HandleFunc("/swagger", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/swagger/swagger.html", http.StatusFound)
	})

	baseUrl := buildBaseURL(cfg)
	repo := getRepo(cfg)

	mux.HandleFunc("POST /api/v1/shorten", addHandler.New(addService.New(baseUrl, repo, &shortener.Hash{})).Handle)
	mux.HandleFunc("GET /api/v1/shorten", shortHandler.New(shortService.New(baseUrl, repo)).Handle)
	mux.HandleFunc("GET /api/v1/original", originalHandler.New(repo).Handle)
	mux.HandleFunc("GET /link/{short_id}", redirectHandler.New(repo).Handle)

	return mux
}

func buildBaseURL(cfg *config.Config) string {
	if cfg.Env.IsLocal() {
		return fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
	}

	return cfg.Host
}

func getRepo(cfg *config.Config) storage.Storage {
	if cfg.StorageType.IsMemory() {
		return cache.NewCache(cfg.Capacity)
	} else {
		client, err := postgresql.NewClient()
		if err != nil {
			log.Fatalf("Postgresql is not running, err: %v", err)
		}
		return db.NewRepository(client)
	}
}
