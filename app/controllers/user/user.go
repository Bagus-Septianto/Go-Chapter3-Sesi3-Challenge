package user

import (
	"challenge294/app/services"

	"gorm.io/gorm"
)

type UserController struct {
	UserService services.IUserService
	DB          *gorm.DB
}

func NewUserController(userService services.IUserService, db *gorm.DB) *UserController {
	return &UserController{
		UserService: userService,
		DB:          db,
	}
}