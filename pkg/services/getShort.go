package services

import "fmt"

type GetShortProvider interface {
	GetShort(link string) (string, bool)
}

type GetShortService struct {
	name              string
	baseUrl           string
	shortenerProvider GetShortProvider
}

func NewGetShortService(baseUrl string, shortenerProvider GetShortProvider) *GetShortService {
	return &GetShortService{
		name:              "get short link service",
		baseUrl:           baseUrl,
		shortenerProvider: shortenerProvider,
	}
}

func (s GetShortService) Get(link string) (string, bool) {
	shortLink, ok := s.shortenerProvider.GetShort(link)
	if !ok {
		return shortLink, ok
	}

	return fmt.Sprintf("%s/link/%s", s.baseUrl, shortLink), ok
}
