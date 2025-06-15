package repositories

import "testing"

func TestLRUEviction(t *testing.T) {
	cache := NewShortenCache(2)
	cache.Add("a", "1")
	cache.Add("b", "2")
	cache.Add("c", "3")

	if _, ok := cache.GetOriginal("a"); ok {
		t.Error("Expected 'a' to be evicted")
	}
}
