package comment

import (
	"github.com/shardent/messec-be/infra/database"
	"github.com/shardent/messec-be/infra/logger"
)

func Create(model interface{}) error {
	err := database.DB.Create(model).Error
	if err != nil {
		logger.Errorf("error, cannot create new comment", err)
	}
	return err
}

func Delete(model interface{}, commentID uint) error {
	err := database.DB.Delete(model, commentID).Error
	if err != nil {
		logger.Errorf("error, deleting comment")
	}

	return err
}
