package storage

type Storage interface {
	Add(shortLink string, originalLink string)
	GetShort(link string) (string, error)
	GetOriginal(shortID string) (string, error)
}
