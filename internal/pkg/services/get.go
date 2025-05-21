package services

type GetProvider interface {
	GetOriginal(shortLink string) (string, bool)
}

type GetService struct {
	name              string
	shortenerProvider GetProvider
}

func NewGetService(shortenerProvider GetProvider) *GetService {
	return &GetService{
		name:              "get short link service",
		shortenerProvider: shortenerProvider,
	}
}

func (s GetService) Get(shortLink string) (string, bool) {
	return s.shortenerProvider.GetOriginal(shortLink)
}
