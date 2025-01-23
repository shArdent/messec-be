package post

import (

	"github.com/shardent/messec-be/infra/database"
	"github.com/shardent/messec-be/infra/logger"
)

func GetPostsByUserId(userID string) ([]Post, error) {
	var posts []Post

	err := database.DB.Preload("Comment").Where("user_id = ?", userID).Order("created_at DESC").Find(&posts).Error
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

func Delete(model interface{}, ID uint) error {
	err := database.DB.Delete(model, ID).Error
	if err != nil {
		logger.Errorf("error, deleting comment")
	}

	return err
}
