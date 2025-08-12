package shortener

import "errors"

var (
	ErrEmptyLink  = errors.New("incorrect Link")
	ErrInvalidURL = errors.New("invalid URL format")
)
