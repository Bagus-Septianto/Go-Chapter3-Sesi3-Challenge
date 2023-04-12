package product

import (
	"challenge294/app/models"

	"gorm.io/gorm"
)

func (repository ProductRepository) ReadAll(db *gorm.DB) (product *[]models.Product, err error) {
	err = db.Debug().Find(&product).Error
	return product, err
}
