package short

import (
	"encoding/json"
	"net/http"

	"github.com/Racuwcka/shorter-url/internal/handler/shortener"
)

type getter interface {
	GetShort(link string) (string, error)
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
	link := r.URL.Query().Get("link")
	req := &shortener.OriginalLinkRequest{
		Link: link,
	}

	if err := req.Validate(); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	shortLink, err := h.g.GetShort(req.Link)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	res := &shortener.ShortLinkResponse{
		ShortLink: shortLink,
	}

	raw, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, _ = w.Write(raw)
}
