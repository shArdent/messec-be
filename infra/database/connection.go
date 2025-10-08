package database

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB  *gorm.DB
	err error
)

func DbConnection(masterDSN string) error {
	db := DB
	logMode := viper.GetBool("DB_LOG_MODE")

	logLevel := logger.Silent
	if logMode {
		logLevel = logger.Info
	}

	db, err = gorm.Open(mysql.Open(masterDSN), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})

	if err != nil {
		log.Printf("Db connection error: %v", err)
		return err
	}
	DB = db
	return nil
}

// GetDB connection
func GetDB() *gorm.DB {
	return DB
}

func TestConnection(dsn string) error {
	testDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	sqlDB, err := testDB.DB()
	if err != nil {
		return err
	}
	defer sqlDB.Close()

	fmt.Printf("aman")

	return sqlDB.Ping()
}
