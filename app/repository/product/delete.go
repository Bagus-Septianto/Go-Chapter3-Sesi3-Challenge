package product

import (
	"challenge294/app/models"
	"errors"

	"gorm.io/gorm"
)

func (repository ProductRepository) Delete(db *gorm.DB, productID int) (err error) {
	dbreturn := db.Where("id = ?", productID).Delete(&models.Product{})
	if dbreturn.Error != nil {
		return dbreturn.Error
	}
	// No RowsAffected (= 0), hacky solution: put if statement here and make errors
	if dbreturn.RowsAffected == 0 {
		return errors.New("data doesn't exist")
	}
	return nil
}
