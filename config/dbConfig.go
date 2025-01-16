package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type DbConfiguration struct {
	Driver   string
	Dbname   string
	Username string
	Password string
	Host     string
	Port     string
	LogMode  bool
}

func DbConfig() (string, string) {
	masterDBName := viper.GetString("MASTER_DB_NAME")
	masterDBUser := viper.GetString("MASTER_DB_USER")
	masterDBPassword := viper.GetString("MASTER_DB_PASSWORD")
	masterDBHost := viper.GetString("MASTER_DB_HOST")
	masterDBPort := viper.GetString("MASTER_DB_PORT")
	// masterDBSslMode := viper.GetString("MASTER_SSL_MODE")

	replicaDBName := viper.GetString("REPLICA_DB_NAME")
	replicaDBUser := viper.GetString("REPLICA_DB_USER")
	replicaDBPassword := viper.GetString("REPLICA_DB_PASSWORD")
	replicaDBHost := viper.GetString("REPLICA_DB_HOST")
	replicaDBPort := viper.GetString("REPLICA_DB_PORT")
	// replicaDBSslMode := viper.GetString("REPLICA_SSL_MODE")

	masterDBSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		masterDBUser, masterDBPassword, masterDBHost, masterDBPort, masterDBName)

	replicaDBSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		replicaDBUser, replicaDBPassword, replicaDBHost, replicaDBPort, replicaDBName)

	return masterDBSN, replicaDBSN
}
