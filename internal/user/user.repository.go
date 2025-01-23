package user

import (
	"github.com/shardent/messec-be/infra/database"
	"github.com/shardent/messec-be/infra/logger"
)

func GetAllUser() ([]UserDto, error) {
	var usersDtos []UserDto

	err := database.DB.Model(&User{}).Find(&usersDtos).Error
	if err != nil {
		logger.Errorf("Failed to get data from db", err)
		return nil, err
	}

	return usersDtos, nil
}
