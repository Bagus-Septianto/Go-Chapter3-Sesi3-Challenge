package product

import (
	"challenge294/app/models"

	"gorm.io/gorm"
)

func (repository ProductRepository) Read(db *gorm.DB, productID int) (product *models.Product, err error) {
	err = db.Debug().First(&product, "id = ?", productID).Error
	return product, err
}
