package test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/bootcamp-go/desafio-cierre-testing/internal/products"
	"github.com/bootcamp-go/desafio-cierre-testing/internal/products/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var (
	LIST_PRODUCT_RETURN = []products.Product{
		{
			ID:          "mock",
			SellerID:    "FEX112AC",
			Description: "generic product",
			Price:       123.55,
		},
	}
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

	testCases := []struct {
		Name               string
		ExpectedStatusCode int
		RepoReturnOnGet    []products.Product
		RepoErrorOnGet     error
		ExpectedData       []products.Product
		Path               string
		ExpectedMessage    string
	}{
		{
			Name:               "Get products happy path return a 200 status code",
			ExpectedStatusCode: http.StatusOK,
			RepoReturnOnGet:    LIST_PRODUCT_RETURN,
			RepoErrorOnGet:     nil,
			ExpectedData:       LIST_PRODUCT_RETURN,
			Path:               "/api/v1/products?seller_id=FEX112AC",
			ExpectedMessage:    "Ok",
		},
		{
			Name:               "Get products returns 400 status code when paremeter is empty",
			ExpectedStatusCode: http.StatusBadRequest,
			RepoReturnOnGet:    nil,
			RepoErrorOnGet:     nil,
			ExpectedData:       nil,
			Path:               "/api/v1/products?seller_id=",
			ExpectedMessage:    "Invalid parameter",
		},
		{
			Name:               "Get products returns 500 status code when repository returns an error",
			ExpectedStatusCode: http.StatusInternalServerError,
			RepoReturnOnGet:    nil,
			RepoErrorOnGet:     errors.New("I'm an unexpected error"),
			ExpectedData:       nil,
			Path:               "/api/v1/products?seller_id=FEX112AC",
			ExpectedMessage:    "Internal server error",
		},
		{
			Name:               "Get products returns an empty data",
			ExpectedStatusCode: http.StatusOK,
			RepoReturnOnGet:    nil,
			RepoErrorOnGet:     nil,
			ExpectedData:       nil,
			Path:               "/api/v1/products?seller_id=OTHERID",
			ExpectedMessage:    "Ok",
		},
	}

	for _, tc := range testCases {

		t.Run(tc.Name, func(t *testing.T) {
			// Arrange.
			repository.Reset()
			expectedStatusCode := tc.ExpectedStatusCode
			repository.ReturnOnGet = tc.RepoReturnOnGet
			repository.ErrorOnGet = tc.RepoErrorOnGet

			expectedData := tc.ExpectedData

			req, res := NewRequest(http.MethodGet, tc.Path, "")

			// Act.
			server.ServeHTTP(res, req)
			var r responseProducts
			err := json.Unmarshal(res.Body.Bytes(), &r)

			// Assert.
			assert.NoError(t, err)
			assert.Equal(t, expectedStatusCode, res.Code)
			assert.Equal(t, expectedData, r.Data)
			assert.Equal(t, tc.ExpectedMessage, r.Message)
		})
	}
}
