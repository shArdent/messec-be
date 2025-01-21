package post

import (
	"github.com/shardent/messec-be/infra/database"
	"github.com/shardent/messec-be/infra/logger"
)

func GetPostsByUserId(userId string) ([]Post, error) {
	var posts []Post
	err := database.DB.Where("user_id = ?", userId).Find(&posts).Error
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func CreatePost(model interface{}) error {
	err := database.DB.Create(model).Error
	if err != nil {
		logger.Errorf("error create new post", err)
	}

	return err
}
