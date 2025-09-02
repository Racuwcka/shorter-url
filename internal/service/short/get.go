package short

import (
	"github.com/Racuwcka/shorter-url/internal/utils/shortlink"
)

type provider interface {
	GetShort(link string) (string, error)
}

type GetShortService struct {
	baseUrl  string
	provider provider
}

func New(baseUrl string, p provider) *GetShortService {
	return &GetShortService{
		baseUrl:  baseUrl,
		provider: p,
	}
}

func (s *GetShortService) GetShort(link string) (string, error) {
	shortID, err := s.provider.GetShort(link)
	if err != nil {
		return shortID, err
	}

	return shortlink.Create(s.baseUrl, shortID), nil
}
