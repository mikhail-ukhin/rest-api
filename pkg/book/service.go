package book

type Service interface {
}

type service struct {
	r Repository
}

func (s *service) NewService(r Repository) Service {
	return &service{r}
}
