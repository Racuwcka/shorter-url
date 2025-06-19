package short

import "fmt"

type provider interface {
	GetShort(link string) (string, error)
}

type GetShortService struct {
	baseUrl string
	p       provider
}

func NewGetShortService(baseUrl string, provider provider) *GetShortService {
	return &GetShortService{
		baseUrl: baseUrl,
		p:       provider,
	}
}

func (s GetShortService) GetShort(link string) (string, error) {
	shortLink, err := s.p.GetShort(link)
	if err != nil {
		return shortLink, err
	}

	return fmt.Sprintf("%s/link/%s", s.baseUrl, shortLink), nil
}
