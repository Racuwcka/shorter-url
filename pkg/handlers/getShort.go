package handlers

import (
	"encoding/json"
	"net/http"
)

type GetShortRequest struct {
	link string
}

type GetShortResponse struct {
	ShortLink string `json:"short_link"`
}

func (r GetShortRequest) Validate() error {
	if r.link == "" {
		return errorEmptyLink
	}
	return nil
}

type GetterShort interface {
	Get(link string) (string, bool)
}

type GetShortHandler struct {
	name string
	GetterShort
}

func NewGetShortHandler(getterShorten GetterShort) *GetShortHandler {
	return &GetShortHandler{
		name:        "get shorten handler",
		GetterShort: getterShorten,
	}
}

func (h GetShortHandler) Handle(w http.ResponseWriter, r *http.Request) {
	link := r.URL.Query().Get("link")
	req := &GetShortRequest{
		link: link,
	}

	if err := req.Validate(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	shortLink, ok := h.Get(req.link)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	res := &GetShortResponse{
		ShortLink: shortLink,
	}

	raw, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, _ = w.Write(raw)
}
