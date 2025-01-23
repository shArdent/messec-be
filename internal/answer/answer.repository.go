package answer

import (
	"github.com/shardent/messec-be/infra/database"
	"github.com/shardent/messec-be/infra/logger"
)

func CreateComment(model interface{}) error {
	err := database.DB.Create(model).Error
	if err != nil {
		logger.Errorf("error, cannot create new comment", err)
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
