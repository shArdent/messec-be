package config

import (
	"github.com/shardent/messec-be/infra/logger"
	"github.com/spf13/viper"
)

type Configuration struct {
	Server   ServerConfiguration
	Database DbConfiguration
}

func SetupConfig() error {
    var configuration *Configuration
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		logger.Errorf("Error reading config file %s", err)
		return err
	}

	err := viper.Unmarshal(&configuration)
    if err != nil {
        logger.Errorf("error to decode, %v", err)
        return err
    }

    return nil
}
