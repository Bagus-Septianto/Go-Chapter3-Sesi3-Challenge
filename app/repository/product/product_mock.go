package product

import (
	"challenge294/app/models"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type ProductRepositoryMock struct {
	Mock mock.Mock
}

func (repository *ProductRepositoryMock) Read(db *gorm.DB, productID int) (product *models.Product, err error) {
	arguments := repository.Mock.Called(productID)

	if arguments.Get(0) == nil {
		return nil, arguments.Error(1)
	}

	product = arguments.Get(0).(*models.Product)

	return product, nil
}

func (repository *ProductRepositoryMock) ReadAll(db *gorm.DB) (product *[]models.Product, err error)  {
	arguments := repository.Mock.Called()

	if arguments.Get(0) == nil {
		return nil, arguments.Error(1)
	}

	product = arguments.Get(0).(*[]models.Product)

	return product, nil
}

// karena yang diminta hanya read dan readall. saya hanya membuat mock untuk read dan readall
func (repository *ProductRepositoryMock) Create(db *gorm.DB, p *models.Product) (product *models.Product, err error) {
	return nil, nil
}
func (repository *ProductRepositoryMock) Update(db *gorm.DB, p *models.Product, productID int) (product *models.Product, err error) {
	return nil, nil
}
func (repository *ProductRepositoryMock) Delete(db *gorm.DB, productID int) (err error) {
	return nil
}
