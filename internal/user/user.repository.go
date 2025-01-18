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

func GetUserByEmailOrEmail(existUser *User, user User) error {
	err := database.DB.Where("email = ?", user.Email).Or("username = ?", user.Username).First(&existUser).Error
	if err != nil {
		logger.Errorf("Error query data ", err.Error)
		return err
	}

	return nil
}
