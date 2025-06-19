package redirect

import (
	"net/http"

	"github.com/Racuwcka/shorter-url/internal/handler/shortener"
)

type getter interface {
	Get(shortLink string) (string, error)
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
	shortID := r.PathValue("shortID")
	req := &shortener.ShortIDRequest{
		ShortID: shortID,
	}

	if err := req.Validate(); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	link, err := h.g.Get(req.ShortID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	http.Redirect(w, r, link, http.StatusFound)
}
