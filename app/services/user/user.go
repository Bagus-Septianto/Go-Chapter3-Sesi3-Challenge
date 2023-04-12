package user
import (
	repo "challenge294/app/repository"

	"gorm.io/gorm"
)

type UserService struct {
	UserRepository repo.IUserRepository
	DB             *gorm.DB
}

func NewUserService(repository repo.IUserRepository, db *gorm.DB) *UserService {
	return &UserService{
		UserRepository: repository,
		DB: db,
	}
}