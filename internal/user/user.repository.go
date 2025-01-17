package user

import (
	"github.com/shardent/messec-be/infra/database"
	"github.com/shardent/messec-be/infra/logger"
)

func CreateNew(model interface{}) error {
	err := database.DB.Create(model).Error
	if err != nil {
		logger.Errorf("error, cannot create new user %v", err)
	}

	return err
}

func GetAllUser() ([]*UserDto, error) {
	var users []User
	var usersDtos []*UserDto

	result := database.DB.Find(&users)
	if result.Error != nil {
		logger.Errorf("Failed to get data from db", result.Error)
		return nil, result.Error
	}

	for _, user := range users {
		usersDtos = append(usersDtos, &UserDto{
			ID:        user.ID,
			Username:  user.Username,
			Name:      user.Name,
			Email:     user.Email,
			Bio:       user.Bio,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}

	return usersDtos, nil
}
