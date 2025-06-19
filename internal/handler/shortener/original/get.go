package original

import (
	"encoding/json"
	"net/http"

	"github.com/Racuwcka/shorter-url/internal/handler/shortener"
)

type getter interface {
	GetOriginal(shortLink string) (string, error)
}

type Handler struct {
	g getter
}

func New(getter getter) *Handler {
	return &Handler{
		g: getter,
	}
}

func (h Handler) Handle(w http.ResponseWriter, r *http.Request) {
	shortLink := r.URL.Query().Get("short_link")
	req := &shortener.ShortLinkRequest{
		ShortLink: shortLink,
	}

	if err := req.Validate(); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	link, err := h.g.GetOriginal(req.ShortLink)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	res := &shortener.OriginalLinkResponse{
		Link: link,
	}

	raw, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, _ = w.Write(raw)
}
