package services

import (
	"crypto/sha256"
	"fmt"
)

type ShortenProvider interface {
	AddShort(shortLink string, link string)
	GetShort(link string) (string, bool)
}

type AddService struct {
	name            string
	shortenProvider ShortenProvider
}

func NewAddService(shortenProvider ShortenProvider) *AddService {
	return &AddService{
		name:            "shorter add service",
		shortenProvider: shortenProvider,
	}
}

func getShortLink(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return fmt.Sprintf("%s/%x", "http://localhost:8080", h.Sum(nil)[:8])
}

func (s AddService) Add(link string) string {
	if value, ok := s.shortenProvider.GetShort(link); ok {
		return value
	}

	shortLink := getShortLink(link)

	s.shortenProvider.AddShort(shortLink, link)

	return shortLink
}
