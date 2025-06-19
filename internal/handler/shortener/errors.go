package shortener

import "errors"

var (
	errEmptyLink  = errors.New("incorrect Link")
	errInvalidURL = errors.New("invalid URL format")
)
