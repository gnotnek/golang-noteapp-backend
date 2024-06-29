package services

import (
	"github.com/gnotnek/golang-noteapp-backend/models"
)

func CreateUser(user models.User) error {
	return DB.Create(&user).Error
}

func FindUserByUsername(username string) (models.User, error) {
	var user models.User
	err := DB.Where("username = ?", username).First(&user).Error
	return user, err
}
