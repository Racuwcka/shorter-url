package services

import (
	"fmt"
)

type GetOriginalProvider interface {
	GetOriginal(shortLink string) (string, bool)
}

type GetOriginalService struct {
	name              string
	baseUrl           string
	shortenerProvider GetOriginalProvider
}

func NewGetOriginalService(baseUrl string, getOriginalProvider GetOriginalProvider) *GetOriginalService {
	return &GetOriginalService{
		name:              "get original link service",
		baseUrl:           baseUrl,
		shortenerProvider: getOriginalProvider,
	}
}

func (s GetOriginalService) GetOriginal(shortLink string) (string, bool) {
	uri := fmt.Sprintf("%s/link/", s.baseUrl)
	shortLink = shortLink[len(uri):]

	return s.shortenerProvider.GetOriginal(shortLink)
}
