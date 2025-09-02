package add

import (
	"github.com/Racuwcka/shorter-url/internal/utils/shortlink"
)

type provider interface {
	Add(shortLink string, link string)
	GetShort(link string) (string, error)
}

type linkShortener interface {
	Generate(link string) string
}

type Service struct {
	baseUrl   string
	provider  provider
	shortener linkShortener
}

func New(baseUrl string, p provider, s linkShortener) *Service {
	return &Service{
		baseUrl:   baseUrl,
		provider:  p,
		shortener: s,
	}
}

func (s *Service) Add(link string) string {
	shortID, err := s.provider.GetShort(link)
	if err != nil {
		shortID = s.shortener.Generate(link)

		s.provider.Add(shortID, link)

		return shortlink.Create(s.baseUrl, shortID)
	}

	return shortlink.Create(s.baseUrl, shortID)
}
