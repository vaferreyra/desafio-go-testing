package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/bootcamp-go/desafio-cierre-testing/internal/products"
	"github.com/bootcamp-go/desafio-cierre-testing/internal/products/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServerForProductsTest(rp products.Repository) *gin.Engine {
	// Instances.
	service := products.NewService(rp)

	// Server.
	server := gin.Default()

	// -> routes
	routes := server.Group("/api/v1")
	{
		h := products.NewHandler(service)
		group := routes.Group("/products")
		group.GET("", h.GetProducts)
	}

	return server
}

func NewRequest(method, path, body string) (req *http.Request, res *httptest.ResponseRecorder) {
	// request
	req = httptest.NewRequest(method, path, strings.NewReader(body))
	req.Header.Add("Content-Type", "application/json")

	// response
	res = httptest.NewRecorder()

	return
}

type responseProducts struct {
	Message string             `json:"message"`
	Data    []products.Product `json:"data"`
}

func TestGetProducts(t *testing.T) {
	repository := mocks.NewFakeRepository()
	server := createServerForProductsTest(repository)
	t.Run("Get products happy path return a 200 status code", func(t *testing.T) {
		// Arrange.
		repository.ReturnOnGet = []products.Product{
			{
				ID:          "mock",
				SellerID:    "FEX112AC",
				Description: "generic product",
				Price:       123.55,
			},
		}

		expectedStatusCode := http.StatusOK
		expectedData := []products.Product{
			{
				ID:          "mock",
				SellerID:    "FEX112AC",
				Description: "generic product",
				Price:       123.55,
			},
		}

		req, res := NewRequest(http.MethodGet, "/api/v1/products?seller_id=FEX112AC", "")

		// Act.
		server.ServeHTTP(res, req)
		var r responseProducts
		err := json.Unmarshal(res.Body.Bytes(), &r)

		// Assert.
		assert.NoError(t, err)
		assert.Equal(t, expectedStatusCode, res.Code)
		assert.Equal(t, expectedData, r.Data)
	})
}
