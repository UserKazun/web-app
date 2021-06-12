package services

import (
	"local.packages/models"
)

func CreateUser(user models.User) (models.User, error) {
	err := db.Create(&user).Error
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
