package products

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetProducts(t *testing.T) {
	repositoryMockTestify := NewRepositoryMockTestify()
	service := NewService(repositoryMockTestify)
	handler := NewHandler(service)

	router := gin.New()
	router.GET("/products", handler.GetProducts)

	t.Run("sucess", func(t *testing.T) {
		//Arange
		IDSellerToSearch := "FEX112AC"
		expectedHTTPStatusCode := 200
		expectedHTTPHeaders := http.Header{
			"Content-Type": []string{"application/json; charset=utf-8"},
		}
		expectedResponse := `[
			{
				"ID": "mock",
				"SellerID": "FEX112AC",
				"Description": "generic product",
				"Price": 123.55
			}
		]`

		responseRecorder := httptest.NewRecorder()

		repositoryMockTestify.On("GetAllBySeller", "FEX112AC").Return(
			[]Product{
				{
					ID:          "mock",
					SellerID:    "FEX112AC",
					Description: "generic product",
					Price:       123.55,
				}},
			nil)

		//Act
		router.ServeHTTP(responseRecorder, httptest.NewRequest(
			http.MethodGet,
			"/products?seller_id="+IDSellerToSearch,
			nil,
		))
		//Assert
		assert.Equal(t, expectedHTTPStatusCode, responseRecorder.Code)
		assert.Equal(t, expectedHTTPHeaders, responseRecorder.HeaderMap)
		assert.JSONEq(t, expectedResponse, responseRecorder.Body.String())
	})

	t.Run("param required", func(t *testing.T) {
		//Arange
		noExistIdSellerError := errors.New("seller_id query param is required")
		IDSellerToSearch := ""
		expectedHTTPStatusCode := 400
		expectedHTTPHeaders := http.Header{
			"Content-Type": []string{"application/json; charset=utf-8"},
		}
		expectedResponse := `{
			"error": "seller_id query param is required"
		}`

		responseRecorder := httptest.NewRecorder()

		repositoryMockTestify.On("GetAllBySeller", "").Return(
			[]Product{},
			noExistIdSellerError)

		//Act
		router.ServeHTTP(responseRecorder, httptest.NewRequest(
			http.MethodGet,
			"/products?seller_id="+IDSellerToSearch,
			nil,
		))
		//Assert
		assert.Equal(t, expectedHTTPStatusCode, responseRecorder.Code)
		assert.Equal(t, expectedHTTPHeaders, responseRecorder.HeaderMap)
		assert.JSONEq(t, expectedResponse, responseRecorder.Body.String())
	})

	t.Run("sucess", func(t *testing.T) {
		//Arange
		noExistIdSellerError := errors.New("error in repository")
		IDSellerToSearch := "F"
		expectedHTTPStatusCode := 500
		expectedHTTPHeaders := http.Header{
			"Content-Type": []string{"application/json; charset=utf-8"},
		}
		expectedResponse := `{
			"error": "error in repository"
		}`

		responseRecorder := httptest.NewRecorder()

		repositoryMockTestify.On("GetAllBySeller", "F").Return(
			[]Product{},
			noExistIdSellerError)

		//Act
		router.ServeHTTP(responseRecorder, httptest.NewRequest(
			http.MethodGet,
			"/products?seller_id="+IDSellerToSearch,
			nil,
		))
		//Assert
		assert.Equal(t, expectedHTTPStatusCode, responseRecorder.Code)
		assert.Equal(t, expectedHTTPHeaders, responseRecorder.HeaderMap)
		assert.JSONEq(t, expectedResponse, responseRecorder.Body.String())
	})
}
