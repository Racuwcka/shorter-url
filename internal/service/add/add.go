package add

import (
	"crypto/sha256"
	"fmt"
)

type provider interface {
	Add(shortLink string, link string)
	GetShort(link string) (string, error)
}

type Service struct {
	baseUrl string
	p       provider
}

func New(baseUrl string, provider provider) *Service {
	return &Service{
		baseUrl: baseUrl,
		p:       provider,
	}
}

func generateShortHash(link string) string {
	h := sha256.Sum256([]byte(link))
	return fmt.Sprintf("%x", h[:6])
}

func (s *Service) Add(link string) string {
	if value, err := s.p.GetShort(link); err == nil {
		return value
	}

	shortLink := generateShortHash(link)

	s.p.Add(shortLink, link)

	return fmt.Sprintf("%s/link/%s", s.baseUrl, shortLink)
}
