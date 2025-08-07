package redirect

import (
	"net/http"

	"github.com/Racuwcka/shorter-url/internal/handler/shortener/dto"
)

type getProvider interface {
	GetOriginal(shortLink string) (string, error)
}

type Handler struct {
	g getProvider
}

func New(getter getProvider) *Handler {
	return &Handler{
		g: getter,
	}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	shortID := r.PathValue("short_id")
	req := &dto.ShortIDRequest{
		ShortID: shortID,
	}

	if err := req.Validate(); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	link, err := h.g.GetOriginal(req.ShortID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	http.Redirect(w, r, link, http.StatusFound)
}
