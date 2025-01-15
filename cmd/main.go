package main

import (
	"time"

	"github.com/shardent/messec-be/config"
	"github.com/shardent/messec-be/infra/database"
	"github.com/shardent/messec-be/infra/logger"
	"github.com/shardent/messec-be/routes"
	"github.com/spf13/viper"
)

func main() {
    viper.SetDefault("SERVER_TIMEZONE", "Asia/Jakarta")
    loc, _ := time.LoadLocation(viper.GetString("SERVER_TIMEZONE"))
    time.Local = loc


    if err := config.SetupConfig();err != nil {
        logger.Fatalf("config SetupConfig() error %s", err)
    }

    masterDSN, replicaDSN := config.DbConfig()

    if err := database.DbConnection(masterDSN, replicaDSN); err != nil {
        logger.Fatalf("database DbConnection error : %s", err)
    }

    migrations.Migrate()
    router := routes.SetupRoutes()

    logger.Fatalf("%v", router.Run(config.ServerConfig()))

}
