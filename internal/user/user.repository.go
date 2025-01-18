package user

import (
	"github.com/shardent/messec-be/infra/database"
	"github.com/shardent/messec-be/infra/logger"
)

func GetAllUser() ([]*UserDto, error) {
	var users []User
	var usersDtos []*UserDto

	err := database.DB.Find(&users).Error
	if err != nil {
		logger.Errorf("Failed to get data from db", err)
		return nil, err
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


