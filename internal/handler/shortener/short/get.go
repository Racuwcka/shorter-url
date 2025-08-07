package short

import (
	"encoding/json"
	"net/http"

	"github.com/Racuwcka/shorter-url/internal/handler/shortener/dto"
)

type getterShortService interface {
	GetShort(link string) (string, error)
}

type Handler struct {
	g getterShortService
}

func New(getter getterShortService) *Handler {
	return &Handler{
		g: getter,
	}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	link := r.URL.Query().Get("link")
	req := &dto.LinkRequest{
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

	res := &dto.ShortLinkResponse{
		ShortLink: shortLink,
	}

	raw, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, _ = w.Write(raw)
}
