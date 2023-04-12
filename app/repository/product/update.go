package product

import (
	"challenge294/app/models"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (repository ProductRepository) Update(db *gorm.DB, p *models.Product, productID int) (product *models.Product, err error) {
	dbreturn := db.Model(&p).Clauses(clause.Returning{}).Where("id = ?", productID).Updates(models.Product{Title: p.Title, Description: p.Description})
	if dbreturn.Error != nil {
		return nil, dbreturn.Error
	}
	// No RowsAffected (= 0), hacky solution: put if statement here and make errors
	if dbreturn.RowsAffected == 0 {
		return nil, errors.New("data doesn't exist")
	}
	return p, nil
}