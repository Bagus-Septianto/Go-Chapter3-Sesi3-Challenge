package product

import (
	"challenge294/app/repository"

	"gorm.io/gorm"
)

type ProductService struct {
	ProductRepository repository.IProductRepository
	DB             *gorm.DB
}

func NewProductService(productRepository repository.IProductRepository, db *gorm.DB) *ProductService {
	return &ProductService{
		ProductRepository: productRepository,
		DB: db,
	}
}