package redirect

type provider interface {
	GetOriginal(shortLink string) (string, error)
}

type Service struct {
	p provider
}

func NewGetService(p provider) *Service {
	return &Service{
		p: p,
	}
}

func (s Service) Get(shortLink string) (string, error) {
	return s.p.GetOriginal(shortLink)
}
