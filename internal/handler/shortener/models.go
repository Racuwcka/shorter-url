package shortener

import "net/url"

type OriginalLinkRequest struct {
	Link string `json:"link"`
}

func (r *OriginalLinkRequest) Validate() error {
	if r.Link == "" {
		return errEmptyLink
	}
	if _, err := url.ParseRequestURI(r.Link); err != nil {
		return errInvalidURL
	}
	return nil
}

type ShortLinkRequest struct {
	ShortLink string `json:"short_link"`
}

func (r *ShortLinkRequest) Validate() error {
	if r.ShortLink == "" {
		return errEmptyLink
	}
	if _, err := url.ParseRequestURI(r.ShortLink); err != nil {
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
