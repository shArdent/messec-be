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

	viper.AutomaticEnv()

	if err := viper.Unmarshal(&configuration); err != nil {
		logger.Errorf("Error decoding config: %v", err)
		return err
	}

	return nil
}
