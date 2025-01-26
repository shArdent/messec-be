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
	result := database.DB.
		Table("user u").
		Select(`
            u.id,
            u.username,
            u.email,
            u.name,
            u.bio,
            COUNT(DISTINCT p.id) AS post_count,
            COUNT(DISTINCT q.id) AS question_count
        `).
		Joins("LEFT JOIN posts p ON u.id = p.user_id").
		Joins("LEFT JOIN questions q ON u.id = q.user_id").
		Where("u.id = ?", ID).
		Group("u.id").
		Scan(&model)

	return result.RowsAffected, result.Error
}

func GetUserByQuery(query string) ([]UserQueryDto, error) {
	var users []UserQueryDto

	formattedQuery := fmt.Sprint("%", query, "%")
	result := database.DB.Model(&User{}).Where("email LIKE ?", formattedQuery).Or("username LIKE ?", formattedQuery).Find(&users)

	fmt.Printf("%s", formattedQuery)

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("No users found matching query %s", query)
	}

	return users, result.Error
}

func UpdateUser(exist, userPayload interface{}) error {
	err := database.DB.Model(exist).Updates(userPayload).Error
	return err
}
