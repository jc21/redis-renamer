package main

import (
	"os"

	"redisrenamer/pkg/config"
	"redisrenamer/pkg/helpers"
	"redisrenamer/pkg/logger"
	"redisrenamer/pkg/migrator"
)

var version string

func main() {
	argConfig := config.GetConfig(&version)
	logger.Init(argConfig)
	logger.Trace("Args: %+v", argConfig)
	host := argConfig.GetHost()

	if err := host.Check(); err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	client := helpers.NewRedisClient(host)
	migrator.DoMigration(client, argConfig.KeyFilter, argConfig.KeyPrefix, argConfig.Unsafe)
}
