package user

import (
	"jokenpo-api/domain/model"
	userRepository "jokenpo-api/infrastructure/repository/user"
)

func FindAll() ([]model.User, error) {
	users, err := userRepository.FindAll()

	if err != nil {
		return nil, err
	}

	return users, nil
}
