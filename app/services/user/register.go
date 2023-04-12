package user

import (
	"challenge294/app/models"
)

func (service UserService) Register(u *models.User) (user *models.User, err error) {
	err = service.UserRepository.Register(service.DB, u)
	if err != nil {
		return nil, err
	}
	user = u
	return user, nil
}
