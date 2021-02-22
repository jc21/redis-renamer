package model

import "fmt"

var version *string

// ArgConfig is the settings for passing arguments to the command
type ArgConfig struct {
	Host      string `arg:"--host,required" help:"redis server hostname"`
	Port      int    `arg:"--port" default:"6379" help:"redis server port"`
	DBIndex   int    `arg:"--db" default:"0" help:"redis server db index"`
	RedisUser string `arg:"--redis-user" help:"redis server auth username"`
	RedisPass string `arg:"--redis-pass" help:"redis server auth password"`
	KeyFilter string `arg:"--filter" default:"*" help:"keys filter string"`
	KeyPrefix string `arg:"--prefix,required" help:"key prefix to prepend"`
	Unsafe    bool   `arg:"-u" help:"Don't bother checking if key already has prefix first"`
	Verbose   bool   `arg:"-v" help:"Print a lot more info"`
}

// SetVersion ...
func SetVersion(ver *string) {
	version = ver
}

// Version ...
func (ArgConfig) Version() string {
	return fmt.Sprintf("v%s", *version)
}

// Description returns a simple description of the command
func (ArgConfig) Description() string {
	return `Redis Renamer will simply rename the matching keys with a given prefix.
`
}

// GetHost returns redis server
func (c *ArgConfig) GetHost() RedisServerConfig {
	return RedisServerConfig{
		Hostname: c.Host,
		Port:     c.Port,
		DBIndex:  c.DBIndex,
		Username: c.RedisUser,
		Password: c.RedisPass,
	}
}

func determineAuth(user, pass string) string {
	auth := "None"
	if pass != "" {
		auth = "Password only"
	}
	if user != "" {
		auth = "Username and Password"
	}
	return auth
}
