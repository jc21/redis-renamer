package migrator

import (
	"fmt"
	"os"
	"strings"

	"redisrenamer/pkg/helpers"
	"redisrenamer/pkg/logger"

	redis "github.com/go-redis/redis/v8"
)

// DoMigration ...
func DoMigration(client *redis.Client, keyFilter, keyPrefix string, unsafe bool) {
	size, err := helpers.GetDBSize(client, keyFilter)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	logger.Info("Found %d keys on host", size)

	if size == 0 {
		logger.Warn("Migration was aborted as host has no keys matching filter")
		os.Exit(1)
	}

	logger.Info("Migration running, each dot is ~1,000 keys")
	logger.Info("Scan will iterate ove all keys but only rename those matching the key prefix")

	var cursor uint64
	var n int

	counter := 0
	skipped := 0
	dots := 0

	for {
		var keys []string
		var scanErr error
		keys, cursor, scanErr = client.Scan(helpers.Ctx, cursor, keyFilter, 1000).Result()
		if scanErr != nil {
			logger.Error("Scan Error: %+v", scanErr)
			os.Exit(1)
		}

		for _, key := range keys {
			if !unsafe {
				// Check that key doesn't already contain the prefix
				if strings.HasPrefix(key, keyPrefix) {
					// move to next key
					skipped++
					continue
				}
			}

			newKey := keyPrefix + key
			if renameErr := client.Rename(helpers.Ctx, key, newKey).Err(); renameErr != nil {
				logger.Error("Could not rename '%s' to '%s': %s", key, newKey, renameErr.Error())
			}

			counter++
		}

		fmt.Print(".")
		dots++
		n += len(keys)
		if dots%10 == 0 {
			fmt.Print(" ")
		}
		if dots%50 == 0 {
			fmt.Print("\n")
		}
		if cursor == 0 {
			break
		}
	}

	fmt.Print("\n")
	logger.Info("Migration completed with %d keys, %d skipped for safety :)", counter, skipped)
	os.Exit(0)
}
