package db

import (
	"context"

	"github.com/Racuwcka/shorter-url/internal/storage"
	"github.com/Racuwcka/shorter-url/pkg/client/postgresql"
)

type Repository struct {
	client postgresql.Client
}

var _ storage.Storage = (*Repository)(nil)

func NewRepository(c postgresql.Client) *Repository {
	return &Repository{
		client: c,
	}
}

func (r *Repository) Add(shortLink string, originalLink string) {
	q := `INSERT INTO public.urls (original_url, short_code) VALUES ($1, $2)`

	r.client.QueryRow(context.Background(), q, originalLink, shortLink)
}

func (r *Repository) GetShort(link string) (string, error) {
	q := `SELECT short_code FROM public.urls WHERE original_url = $1`

	var shortID string
	err := r.client.QueryRow(context.Background(), q, link).Scan(&shortID)
	if err != nil {
		return "", err
	}
	return shortID, nil
}

func (r *Repository) GetOriginal(shortID string) (string, error) {
	q := `SELECT original_url FROM public.urls WHERE short_code = $1`

	var link string
	err := r.client.QueryRow(context.Background(), q, shortID).Scan(&link)
	if err != nil {
		return "", err
	}
	return link, nil
}
