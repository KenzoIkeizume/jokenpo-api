package user

import (
	"jokenpo-api/domain/model"
	"jokenpo-api/infrastructure/datastore"
)

type userRepository struct{}

type UserRepository interface {
	FindAll() ([]model.User, error)
}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (us userRepository) FindAll() ([]model.User, error) {
	users, err := datastore.Find()

	if err != nil {
		return nil, err
	}

	return users, nil
}
