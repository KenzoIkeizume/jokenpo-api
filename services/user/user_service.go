package user

import (
	"jokenpo-api/domain/model"
	user_repository "jokenpo-api/infrastructure/repository/user"
)

type userService struct {
	userRepository user_repository.UserRepository
}

type UserService interface {
	FindAll(us []*model.User) ([]*model.User, error)
}

func NewUserService(ur user_repository.UserRepository) UserService {
	return &userService{ur}
}

func (us userService) FindAll(u []*model.User) ([]*model.User, error) {
	users, err := us.userRepository.FindAll(u)

	if err != nil {
		return nil, err
	}

	return users, nil
}
