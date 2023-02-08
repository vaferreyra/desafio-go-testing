package test

import (
	"testing"

	"github.com/bootcamp-go/desafio-cierre-testing/internal/products"
	"github.com/bootcamp-go/desafio-cierre-testing/internal/products/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGetAllBySeller(t *testing.T) {
	t.Run("Get all by seller happy path", func(t *testing.T) {
		// Arrange.
		repository := mocks.NewFakeRepository()
		repository.ReturnOnGet = []products.Product{
			{
				ID:          "mock",
				SellerID:    "FEX112AC",
				Description: "generic product",
				Price:       123.55,
			},
		}

		service := products.NewService(repository)

		expectedResult := []products.Product{
			{
				ID:          "mock",
				SellerID:    "FEX112AC",
				Description: "generic product",
				Price:       123.55,
			},
		}

		// Act.
		obtainedResult, err := service.GetAllBySeller("FEX112AC")

		// Assert.
		assert.NoError(t, err)
		assert.Equal(t, obtainedResult, expectedResult)
	})
}
