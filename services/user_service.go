package service

import (
	"jokenpo-api/domain/model"
	"jokenpo-api/infrastructure/repository"
)

func FindAll() ([]model.User, error) {
	users, err := repository.FindAll()

	if err != nil {
		return nil, err
	}

	return users, nil
}
