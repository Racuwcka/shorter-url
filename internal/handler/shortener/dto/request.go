package dto

import (
	"net/url"

	"github.com/Racuwcka/shorter-url/internal/handler/shortener"
)

type LinkRequest struct {
	Link string `json:"link"`
}

func (r *LinkRequest) Validate() error {
	if r.Link == "" {
		return shortener.ErrEmptyLink
	}
	if _, err := url.ParseRequestURI(r.Link); err != nil {
		return shortener.ErrInvalidURL
	}
	return nil
}

type ShortIDRequest struct {
	ShortID string `json:"short_id"`
}

func (r *ShortIDRequest) Validate() error {
	if r.ShortID == "" {
		return shortener.ErrEmptyLink
	}
	return nil
}

type OriginalLinkResponse struct {
	Link string `json:"link"`
}

type ShortLinkResponse struct {
	ShortLink string `json:"short_link"`
}
