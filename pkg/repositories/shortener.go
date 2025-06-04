package repositories

import (
	"sync"
)

type ShortenCache struct {
	shortToOriginal sync.Map
	originalToShort sync.Map
}

func NewShortenCache() *ShortenCache {
	return &ShortenCache{}
}

func (s *ShortenCache) Add(shortLink string, link string) {
	s.shortToOriginal.Store(shortLink, link)
	s.originalToShort.Store(link, shortLink)
}

func (s *ShortenCache) GetShort(link string) (string, bool) {
	if val, ok := s.originalToShort.Load(link); ok {
		return val.(string), true
	}

	return "", false
}

func (s *ShortenCache) GetOriginal(shortLink string) (string, bool) {
	if val, ok := s.shortToOriginal.Load(shortLink); ok {
		return val.(string), true
	}

	return "", false
}
