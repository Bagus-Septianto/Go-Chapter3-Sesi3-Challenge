package user

import (
	"challenge294/app/models"

	"gorm.io/gorm"
)

func (repository UserRepository) Login(db *gorm.DB, username string) (user *models.User, err error) {
	err = db.Debug().Where("username = ?", username).Take(&user).Error
	return user, err
}
