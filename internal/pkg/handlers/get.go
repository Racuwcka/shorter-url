package handlers

import (
	"net/http"
)

type GetRequest struct {
	shortLink string
}

func (r GetRequest) Validate() error {
	if r.shortLink == "" {
		return errorEmptyLink
	}
	return nil
}

type Getter interface {
	Get(shortLink string) (string, bool)
}

type GetHandler struct {
	name string
	Getter
}

func NewGetHandler(getter Getter) *GetHandler {
	return &GetHandler{
		name:   "shortener get handler",
		Getter: getter,
	}
}

func (h GetHandler) Handle(w http.ResponseWriter, r *http.Request) {
	shortLink := r.PathValue("shortID")
	res := &GetRequest{
		shortLink: shortLink,
	}

	if err := res.Validate(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	link, ok := h.Getter.Get(res.shortLink)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	http.Redirect(w, r, link, http.StatusFound)
}
