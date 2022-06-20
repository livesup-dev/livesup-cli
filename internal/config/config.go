package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func Token() string {
	return getSetting("LIVESUP_TOKEN")
}

func URL() string {
	return getSetting("LIVESUP_URL")
}

func getSetting(key string) string {
	return viper.GetString(key)
}

func Init() {
	viper.SetConfigFile(".env") // optionally look for config in the working directory
	viper.SetDefault("LIVESUP_URL", "http://host.docker.internal:4000/")
	err := viper.ReadInConfig() // Find and read the config file

	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
