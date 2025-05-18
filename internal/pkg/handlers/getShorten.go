package handlers

import (
	"encoding/json"
	"net/http"
)

type GetShortenRequest struct {
	Link string
}

type GetShortenResponse struct {
	ShortLink string `json:"short_link"`
}

func (r GetShortenRequest) Validate() error {
	if r.Link == "" {
		return errorEmptyLink
	}
	return nil
}

type GetterShorten interface {
	Get(link string) string
}

type GetShortenHandler struct {
	name string
	GetterShorten
}

func NewGetShortenHandler(getterShorten GetterShorten) *GetShortenHandler {
	return &GetShortenHandler{
		name:          "get shorten handler",
		GetterShorten: getterShorten,
	}
}

func (h GetShortenHandler) Handle(w http.ResponseWriter, r *http.Request) {
	req := &GetShortenRequest{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := req.Validate(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	shortLink := h.Get(req.Link)

	res := &GetShortenResponse{
		ShortLink: shortLink,
	}

	raw, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, _ = w.Write(raw)
}
