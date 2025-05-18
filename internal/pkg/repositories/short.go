package repositories

import (
	"log"
	"sync"
)

type ShortenCache struct {
	shortToOriginal sync.Map
	originalToShort sync.Map
}

func NewShortenCache() *ShortenCache {
	return &ShortenCache{}
}

func (s *ShortenCache) AddShort(shortLink string, link string) {
	s.shortToOriginal.Store(shortLink, link)
	s.originalToShort.Store(link, shortLink)
}

func (s *ShortenCache) GetShort(link string) (string, bool) {
	if val, ok := s.originalToShort.Load(link); ok {
		log.Println(val)

		return val.(string), true
	}

	return "", false
}
