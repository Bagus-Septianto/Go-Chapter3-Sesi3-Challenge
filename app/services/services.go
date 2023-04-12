package services

import (
	"challenge294/app/models"
)

type IUserService interface {
	Register(u *models.User) (user *models.User, err error)
	Login(username, password string) (token string, err error)
}

type IProductService interface {
	Create(p *models.Product, userID uint) (product *models.Product, err error)
	Read(productID int) (product *models.Product, err error)
	Update(p *models.Product, productID int) (product *models.Product, err error)
	Delete(productID int) (err error)
	ReadAll() (product *[]models.Product, err error)
}
