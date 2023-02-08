package mocks

import "github.com/bootcamp-go/desafio-cierre-testing/internal/products"

type fakeRepository struct {
	ReturnOnGet []products.Product
	ErrorOnGet  error
}

func NewFakeRepository() *fakeRepository {
	return &fakeRepository{}
}

func (fr *fakeRepository) GetAllBySeller(sellerID string) ([]products.Product, error) {
	return fr.ReturnOnGet, fr.ErrorOnGet
}
