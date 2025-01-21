package migrations

import (
	"github.com/shardent/messec-be/infra/database"
	"github.com/shardent/messec-be/infra/logger"
	"github.com/shardent/messec-be/internal/comment"
	"github.com/shardent/messec-be/internal/post"
	"github.com/shardent/messec-be/internal/question"
	"github.com/shardent/messec-be/internal/user"
)

func Migrate() {
	migrationsModels := []interface{}{&user.User{}, &post.Post{}, &comment.Comment{}, &question.Question{}}

	if database.DB == nil {
		logger.Warnf("Database connection is nil")
		return
	}

	err := database.DB.AutoMigrate(migrationsModels...)
	if err != nil {
		logger.Warnf("Failed to do migrations: %v", err)
		return
	}
	logger.Infof("Migrations completed successfully")
}
