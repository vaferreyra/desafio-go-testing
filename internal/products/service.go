package products

import "errors"

var (
	ErrInternalServer = errors.New("Something internal has wrong")
)

type Service interface {
	GetAllBySeller(sellerID string) ([]Product, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetAllBySeller(sellerID string) ([]Product, error) {
	data, err := s.repo.GetAllBySeller(sellerID)
	if err != nil {
		return nil, ErrInternalServer
	}
	return data, err
}
