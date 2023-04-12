package repository

import (
	"challenge294/app/models"

	"gorm.io/gorm"
)

type IUserRepository interface {
	Register(db *gorm.DB, user *models.User) (err error)
	Login(db *gorm.DB, username string) (user *models.User, err error)
}

type IProductRepository interface {
	Create(db *gorm.DB, p *models.Product) (product *models.Product, err error)
	Read(db *gorm.DB, productID int) (product *models.Product, err error)
	Update(db *gorm.DB, p *models.Product, productID int) (product *models.Product, err error)
	Delete(db *gorm.DB, productID int) (err error)
	ReadAll(db *gorm.DB) (product *[]models.Product, err error)
}
