package products

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ------------------------------------------------------------
// stub tests
func TestService_GetAllBySeller(t *testing.T) {
	t.Run("succeed", func(t *testing.T) {
		//Arrange
		repositoryMockTestify := NewRepositoryMockTestify()

		productListExpected := []Product{{
			ID:          "1",
			SellerID:    "hola",
			Description: "saludar",
			Price:       23.5,
		}}

		repositoryMockTestify.On(
			"GetAllBySeller", "hola").Return(productListExpected, nil)

		ser := NewService(repositoryMockTestify)

		//Act
		data, err := ser.GetAllBySeller("hola")

		//Assert
		assert.NoError(t, err)
		assert.Equal(t, productListExpected, data)
	})

	t.Run("error in repository", func(t *testing.T) {
		//Arrange
		noExistIdSellerError := errors.New("no existe id")
		var prodList []Product

		repositoryMockTestify := NewRepositoryMockTestify()

		repositoryMockTestify.On(
			"GetAllBySeller", "hola").Return(prodList, noExistIdSellerError)

		ser := NewService(repositoryMockTestify)

		//Act
		_, err := ser.GetAllBySeller("hola")

		//Assert
		assert.Error(t, err)
	})
}
