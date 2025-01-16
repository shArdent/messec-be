package database

import (
	"github.com/shardent/messec-be/internal/user"
)

func Migrate() {
	migrationsModels := []interface{}{&user.User{}}

	err := DB.AutoMigrate(migrationsModels...)
	if err != nil {
		return
	}
}
