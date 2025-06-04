package handlers

import (
	"encoding/json"
	"net/http"
)

type getOriginalRequest struct {
	shortLink string
}

type getOriginalResponse struct {
	Link string `json:"link"`
}

func (r getOriginalRequest) Validate() error {
	if r.shortLink == "" {
		return errorEmptyLink
	}
	return nil
}

type GetterOriginal interface {
	GetOriginal(shortLink string) (string, bool)
}

type GetOriginalHandler struct {
	name string
	GetterOriginal
}

func NewGetOriginalHandler(getterOriginal GetterOriginal) *GetOriginalHandler {
	return &GetOriginalHandler{
		name:           "get original handler",
		GetterOriginal: getterOriginal,
	}
}

func (h GetOriginalHandler) Handle(w http.ResponseWriter, r *http.Request) {
	shortLink := r.URL.Query().Get("short_link")
	req := &getOriginalRequest{
		shortLink: shortLink,
	}

	if err := req.Validate(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	link, ok := h.GetOriginal(req.shortLink)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	res := &getOriginalResponse{
		Link: link,
	}

	raw, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, _ = w.Write(raw)
}
