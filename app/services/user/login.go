package user

import (
	"challenge294/pkg/helpers"
)

func (service UserService) Login(username, password string) (token string, err error) {
	user, err := service.UserRepository.Login(service.DB, username)
	if (err != nil || !helpers.ComparePass([]byte(user.Password), []byte(password))) {
		return "", err
	}
	token = helpers.GenerateToken(user.ID, user.Username, user.Role)

	return token, nil
}
