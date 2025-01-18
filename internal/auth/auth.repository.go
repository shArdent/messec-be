package auth

import (
	"github.com/shardent/messec-be/infra/database"
	"github.com/shardent/messec-be/infra/logger"
	"github.com/shardent/messec-be/internal/user"
)

func CreateNew(model interface{}) error {
	err := database.DB.Create(model).Error
	if err != nil {
		logger.Errorf("error, cannot create new user %v", err)
	}

	return err
}

func GetUserByEmailOrUsername(existUser, user *user.User) error {
	err := database.DB.Where("email = ?", user.Email).Or("username = ?", user.Username).First(&existUser).Error
	if err != nil {
		logger.Errorf("Error query data ", err.Error)
		return err
	}

	return nil
}
