package service

import (
	"local.packages/model"
)

func CreateUser(user model.User) (model.User, error) {
	err := db.Create(&user).Error
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
