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
