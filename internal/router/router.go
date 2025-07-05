package router

import (
	"fmt"
	"net/http"

	"github.com/Racuwcka/shorter-url/internal/config"
	"github.com/Racuwcka/shorter-url/internal/handler/shortener/add"
	"github.com/Racuwcka/shorter-url/internal/handler/shortener/original"
	"github.com/Racuwcka/shorter-url/internal/handler/shortener/redirect"
	"github.com/Racuwcka/shorter-url/internal/handler/shortener/short"
	addService "github.com/Racuwcka/shorter-url/internal/service/add"
	originalService "github.com/Racuwcka/shorter-url/internal/service/original"
	shortService "github.com/Racuwcka/shorter-url/internal/service/short"
	"github.com/Racuwcka/shorter-url/internal/storage/cache"
)

func New(cfg *config.Config) http.Handler {
	mux := http.NewServeMux()

	baseUrl := buildBaseURL(cfg)

	repo := cache.NewShortenCache(cfg.Capacity)

	mux.Handle("/swagger/", http.StripPrefix("/swagger/", http.FileServer(http.Dir("./../../swagger-ui"))))
	mux.HandleFunc("/swagger", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/swagger/swagger.html", http.StatusFound)
	})

	mux.HandleFunc("POST /api/v1/shorten", add.New(addService.New(baseUrl, repo)).Handle)
	mux.HandleFunc("GET /api/v1/shorten", short.New(shortService.NewGetShortService(baseUrl, repo)).Handle)
	mux.HandleFunc("GET /api/v1/original", original.New(originalService.NewGetOriginalService(baseUrl, repo)).Handle)
	mux.HandleFunc("GET /link/{shortID}", redirect.New(repo).Handle)

	return mux
}

func buildBaseURL(cfg *config.Config) string {
	switch cfg.Env {
	case "local":
		return fmt.Sprintf("http://%s", cfg.Addr)
	default:
		return fmt.Sprintf("https://%s", cfg.Addr)
	}
}
