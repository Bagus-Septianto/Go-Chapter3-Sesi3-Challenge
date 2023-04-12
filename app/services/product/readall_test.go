package product

import (
	"challenge294/app/models"
	"errors"
	"fmt"

	// "fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadAllFound(t *testing.T) {
	products := &[]models.Product{
		{Title: "Title testify0", Description: "Description testify0", UserID: 1},
		{Title: "Title testify1", Description: "Description testify1", UserID: 1},
	}

	productRepository.Mock.On("ReadAll").Return(products, nil).Once()

	results, err := productService.ReadAll()

	assert.Nil(t, err)
	assert.NotNil(t, &results)

	assert.Equal(t, products, results, "result has to be a products data with id '1'")

	for i, result := range *results {
		assert.Equal(t, (*products)[i].Title, result.Title, fmt.Sprintf("result has to be 'Title testify%d'", i))
		assert.Equal(t, (*products)[i].Description, result.Description, fmt.Sprintf("result has to be 'Description testify%d'", i))
	}
}

func TestReadAllNotFound(t *testing.T) {
	// kena di middleware, harusnya arg[0] dari return "Data Not Found"
	productRepository.Mock.On("ReadAll").Return(nil, errors.New("data doesn't exist")).Once()

	result, err := productService.ReadAll() // tapi balikan dari read ini *product dan error

	assert.Nil(t, result)

	assert.NotNil(t, err)
	assert.Equal(t, "data doesn't exist", err.Error(), "error response has to be 'data doesn't exist'")
}
