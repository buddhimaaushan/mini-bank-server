package utils

import "github.com/spf13/viper"

// Config is the configuration for the application.
type Config struct {
	DatabaseURL   string `mapstructure:"DATABASE_URL"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

func LoadConfig(path string) (config Config, err error) {
	// Set default values
	viper.SetDefault("ServerAddress", "0.0.0.0:8080")

	// Add config path and file name
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	// Read in environment variables
	viper.AutomaticEnv()

	// Read in the config file
	err = viper.ReadInConfig()

	if err != nil {
		return config, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return config, err
	}

	return config, err
}
