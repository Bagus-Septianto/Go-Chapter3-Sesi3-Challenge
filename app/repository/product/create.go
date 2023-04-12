package product

import (
	"challenge294/app/models"

	"gorm.io/gorm"
)

func (repository ProductRepository) Create(db *gorm.DB, p *models.Product) (product *models.Product, err error) {
	err = db.Debug().Create(&p).Error
	product = p
	return product, err
}