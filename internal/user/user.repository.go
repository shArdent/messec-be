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


