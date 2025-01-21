package post

import "github.com/shardent/messec-be/infra/database"

func GetPostsByUserId(userId string) ([]Post, error) {
	var posts []Post
	err := database.DB.Where("user_id = ?", userId).Find(&posts).Error
	if err != nil {
		return nil, err
	}

	return posts, nil
}
