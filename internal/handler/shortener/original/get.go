package original

import (
	"encoding/json"
	"net/http"

	"github.com/Racuwcka/shorter-url/internal/handler/shortener/dto"
	"github.com/Racuwcka/shorter-url/internal/utils/shortid"
)

type provider interface {
	GetOriginal(shortID string) (string, error)
}

type Handler struct {
	provider provider
}

func New(g provider) *Handler {
	return &Handler{
		provider: g,
	}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	shortLink := r.URL.Query().Get("link")

	shortID, err := shortid.Get(shortLink)
	if err != nil {
		http.Error(w, "short id not found", http.StatusBadRequest)
		return
	}

	req := &dto.ShortIDRequest{
		ShortID: shortID,
	}

	if err = req.Validate(); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	link, err := h.provider.GetOriginal(req.ShortID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	res := &dto.OriginalLinkResponse{
		Link: link,
	}

	raw, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, _ = w.Write(raw)
}
