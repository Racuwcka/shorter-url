package cache

import (
	"container/list"
	"errors"
	"sync"
)

var (
	errNotFoundOriginalLink = errors.New("not found original shortener")
	errNotFoundShortLink    = errors.New("not found short shortener")
)

type ShortenerCache struct {
	mu              sync.RWMutex
	capacity        int
	shortToOriginal map[string]*list.Element
	originalToShort map[string]*list.Element
	queue           *list.List
}

type cacheItem struct {
	shortKey    string
	originalKey string
}

func NewShortenCache(cap int) *ShortenerCache {
	if cap <= 0 {
		cap = 1000
	}

	return &ShortenerCache{
		capacity:        cap,
		shortToOriginal: make(map[string]*list.Element),
		originalToShort: make(map[string]*list.Element),
		queue:           list.New(),
	}
}

func (s *ShortenerCache) Add(shortLink string, originalLink string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if elem, exists := s.originalToShort[originalLink]; exists {
		s.queue.MoveToFront(elem)
		return
	}

	if elem, exists := s.shortToOriginal[shortLink]; exists {
		s.queue.MoveToFront(elem)
		return
	}

	if s.queue.Len() == s.capacity {
		s.purge()
	}

	item := &cacheItem{
		shortKey:    shortLink,
		originalKey: originalLink,
	}

	elem := s.queue.PushFront(item)
	s.shortToOriginal[shortLink] = elem
	s.originalToShort[originalLink] = elem
}

func (s *ShortenerCache) GetShort(link string) (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if elem, exists := s.originalToShort[link]; exists {
		go s.updateLRU(elem)
		return elem.Value.(*cacheItem).shortKey, nil
	}

	return "", errNotFoundShortLink
}

func (s *ShortenerCache) GetOriginal(shortLink string) (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if elem, exists := s.shortToOriginal[shortLink]; exists {
		go s.updateLRU(elem)
		return elem.Value.(*cacheItem).originalKey, nil
	}

	return "", errNotFoundOriginalLink
}

func (s *ShortenerCache) purge() {
	if elem := s.queue.Back(); elem != nil {
		item := s.queue.Remove(elem).(*cacheItem)
		delete(s.shortToOriginal, item.shortKey)
		delete(s.originalToShort, item.originalKey)
	}
}

func (s *ShortenerCache) updateLRU(elem *list.Element) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.queue.MoveToFront(elem)
}
