package migrations

import (
	"github.com/shardent/messec-be/infra/database"
	"github.com/shardent/messec-be/infra/logger"
	"github.com/shardent/messec-be/internal/user"
)

func Migrate() {
	migrationsModels := []interface{}{&user.User{}}

	err := database.DB.AutoMigrate(migrationsModels...)
	if err != nil {
		logger.Warnf("Failed to do migrations", err)
		return
	}
}
