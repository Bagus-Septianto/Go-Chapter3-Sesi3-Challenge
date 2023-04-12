package product

import (
	"challenge294/app/models"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFound(t *testing.T) {
	product := models.Product{
		Title: "Title testify",
		Description: "Description testify",
		UserID: 1,
	}

	productRepository.Mock.On("Read", 1).Return(&product, nil)

	result, err := productService.Read(1)

	assert.Nil(t, err)

	assert.NotNil(t, result)

	assert.Equal(t, product.Title, result.Title, "result has to be 'Title testify'")
	assert.Equal(t, product.Description, result.Description, "result has to be 'Description testify'")
	assert.Equal(t, &product, result, "result has to be a product data with id '1'")
}

func TestReadNotFound(t *testing.T) {
	// kena di middleware, harusnya arg[0] dari return "Data Not Found"
	productRepository.Mock.On("Read", 2).Return(nil, errors.New("data doesn't exist"))

	result, err := productService.Read(2) // tapi balikan dari read ini *product dan error

	assert.Nil(t, result)
	
	assert.NotNil(t, err)
	assert.Equal(t, "data doesn't exist", err.Error(), "error response has to be 'data doesn't exist'")
}
