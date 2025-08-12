package add

import (
	"encoding/json"
	"net/http"

	"github.com/Racuwcka/shorter-url/internal/handler/shortener/dto"
)

type adderService interface {
	Add(link string) string
}

type Handler struct {
	adderService adderService
}

func New(a adderService) *Handler {
	return &Handler{
		adderService: a,
	}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	req := &dto.LinkRequest{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := req.Validate(); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	shortLink := h.adderService.Add(req.Link)

	res := &dto.ShortLinkResponse{
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
