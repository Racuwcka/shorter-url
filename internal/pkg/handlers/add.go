package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
)

var (
	errorEmptyLink = errors.New("incorrect Link")
)

type AddRequest struct {
	Link string
}

type AddResponse struct {
	ShortLink string `json:"short_link"`
}

func (r AddRequest) Validate() error {
	if r.Link == "" {
		return errorEmptyLink
	}
	return nil
}

type Adder interface {
	Add(link string) string
}

type AddHandler struct {
	name string
	Adder
}

func NewHandler(shorterAdder Adder) *AddHandler {
	return &AddHandler{
		name:  "short Link add handler",
		Adder: shorterAdder,
	}
}

func (h AddHandler) Handle(w http.ResponseWriter, r *http.Request) {
	req := &AddRequest{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := req.Validate(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	shortLink := h.Add(req.Link)

	res := &AddResponse{
		ShortLink: shortLink,
	}

	raw, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, _ = w.Write(raw)
}
