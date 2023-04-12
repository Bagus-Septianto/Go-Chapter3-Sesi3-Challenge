package product

import (
	"challenge294/app/repository/product"

	"github.com/stretchr/testify/mock"
)

var productRepository = &product.ProductRepositoryMock{Mock: mock.Mock{}}
var productService = ProductService{ProductRepository: productRepository}
