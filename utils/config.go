package utils

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config is the configuration for the application.
type Config struct {
	DatabaseURL   string `mapstructure:"DATABASE_URL"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

func LoadConfig(path string) (config Config, err error) {

	// Add config path and file name
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	// Read in environment variables
	viper.BindEnv("DATABASE_URL")
	viper.BindEnv("SERVER_ADDRESS")
	viper.AutomaticEnv()

	// Read in the config file
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			fmt.Println("Config file not found")
		} else {
			// Config file was found but another error was produced
			return config, err
		}
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return config, err
	}
	return config, err
}
