package shortener

import "net/url"

type LinkRequest struct {
	Link string `json:"link"`
}

func (r *LinkRequest) Validate() error {
	if r.Link == "" {
		return errEmptyLink
	}
	if _, err := url.ParseRequestURI(r.Link); err != nil {
		return errInvalidURL
	}
	return nil
}

type ShortIDRequest struct {
	ShortID string
}

func (r *ShortIDRequest) Validate() error {
	if r.ShortID == "" {
		return errEmptyLink
	}
	return nil
}

type OriginalLinkResponse struct {
	Link string `json:"link"`
}

type ShortLinkResponse struct {
	ShortLink string `json:"short_link"`
}
