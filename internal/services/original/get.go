package original

import (
	"fmt"
)

type provider interface {
	GetOriginal(shortLink string) (string, error)
}

type GetOriginalService struct {
	baseUrl string
	p       provider
}

func NewGetOriginalService(baseUrl string, provider provider) *GetOriginalService {
	return &GetOriginalService{
		baseUrl: baseUrl,
		p:       provider,
	}
}

func (s GetOriginalService) GetOriginal(shortLink string) (string, error) {
	uri := fmt.Sprintf("%s/link/", s.baseUrl)
	shortLink = shortLink[len(uri):]

	return s.p.GetOriginal(shortLink)
}
