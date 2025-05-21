package services

import (
	"crypto/sha256"
	"fmt"
)

type AddProvider interface {
	Add(shortLink string, link string)
	GetShort(link string) (string, bool)
}

type AddService struct {
	name              string
	shortenerProvider AddProvider
}

func NewAddService(shortenerProvider AddProvider) *AddService {
	return &AddService{
		name:              "shorter add service",
		shortenerProvider: shortenerProvider,
	}
}

func getShortLink(link string) string {
	h := sha256.Sum256([]byte(link))
	return fmt.Sprintf("%x", h[:6])
}

func (s AddService) Add(link string) string {
	if value, ok := s.shortenerProvider.GetShort(link); ok {
		return value
	}

	shortLink := getShortLink(link)

	s.shortenerProvider.Add(shortLink, link)

	return shortLink
}
