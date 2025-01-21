package question

import (
	"github.com/shardent/messec-be/infra/database"
	"github.com/shardent/messec-be/infra/logger"
)

func Create(model interface{}) error {
	err := database.DB.Create(model).Error
	if err != nil {
		logger.Errorf("Error creating question", model, err.Error())
	}

	return err
}

func GetQuestionByUserId(userId string) (*[]Question, error) {
	var questions *[]Question
	err := database.DB.Where("user_id = ?", userId).Find(&questions).Error
	if err != nil {
		return nil, err
	}

	return questions, nil
}
