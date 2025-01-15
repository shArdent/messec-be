package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type DbConfig struct {
	Driver   string
	Dbname   string
	Username string
	Password string
	Host     string
	Port     string
	LogMode  bool
}

func DbConfiguration() (string, string) {
	masterDBName := viper.GetString("MASTER_DB_NAME")
	masterDBUser := viper.GetString("MASTER_DB_USER")
	masterDBPassword := viper.GetString("MASTER_DB_PASSWORD")
	masterDBHost := viper.GetString("MASTER_DB_HOST")
	masterDBPort := viper.GetString("MASTER_DB_PORT")
	masterDBSslMode := viper.GetString("MASTER_SSL_MODE")

	replicaDBName := viper.GetString("REPLICA_DB_NAME")
	replicaDBUser := viper.GetString("REPLICA_DB_USER")
	replicaDBPassword := viper.GetString("REPLICA_DB_PASSWORD")
	replicaDBHost := viper.GetString("REPLICA_DB_HOST")
	replicaDBPort := viper.GetString("REPLICA_DB_PORT")
	replicaDBSslMode := viper.GetString("REPLICA_SSL_MODE")

	masterDBSN := fmt.Sprintf("host=%s user=%s password=%s dbname=%s prot=%s sslmode=%s", masterDBHost, masterDBUser, masterDBName, masterDBPassword, masterDBPort, masterDBSslMode)

	replicaDBSN := fmt.Sprintf("host=%s user=%s password=%s dbname=%s prot=%s sslmode=%s", replicaDBHost, replicaDBUser, replicaDBName, replicaDBPassword, replicaDBPort, replicaDBSslMode)

	return masterDBSN, replicaDBSN
}
