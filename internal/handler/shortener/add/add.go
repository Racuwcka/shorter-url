package add

import (
	"encoding/json"
	"net/http"

	"github.com/Racuwcka/shorter-url/internal/handler/shortener"
)

type adderService interface {
	Add(link string) string
}

type Handler struct {
	a adderService
}

func New(adder adderService) *Handler {
	return &Handler{
		a: adder,
	}
}

func (h Handler) Handle(w http.ResponseWriter, r *http.Request) {
	req := &shortener.LinkRequest{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := req.Validate(); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	shortLink := h.a.Add(req.Link)

	res := &shortener.ShortLinkResponse{
		ShortLink: shortLink,
	}

	raw, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(raw)
}
