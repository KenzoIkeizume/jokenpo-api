package datastore

import (
	"jokenpo-api/domain/model"
	"time"
)

var users = []model.User{
	model.User{
		ID:        1,
		Age:       10,
		Name:      "kenzo",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
}

func Find() ([]model.User, error) {
	return users, nil
}
