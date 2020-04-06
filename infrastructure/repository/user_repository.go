package repository

import (
	"jokenpo-api/domain/model"
	"jokenpo-api/infrastructure/datastore"
)

func FindAll() ([]model.User, error) {
	users, err := datastore.Find()

	if err != nil {
		return nil, err
	}

	return users, nil
}
