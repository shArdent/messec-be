package user

import (
	"fmt"

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

func GetUser(model interface{}, ID uint) (int64, error) {
	result := database.DB.Model(&User{}).First(&model, ID)

	return result.RowsAffected, result.Error
}

func GetUserByQuery(query string) ([]UserDto, error) {
	var users []UserDto

	formattedQuery := fmt.Sprint("%", query, "%")
	result := database.DB.Model(&User{}).Where("email LIKE ?", formattedQuery).Or("username LIKE ?", formattedQuery).Find(&users)

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("No users found matching query %s", query)
	}

	return users, result.Error
}
