// Package config contains the configuration for the application.
package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type DBMongoConfig struct {
	URI    string `mapstructure:"uri"`
	DBName string `mapstructure:"db_name"`
}

type Config struct {
	Port    string        `mapstructure:"port"`
	DBMongo DBMongoConfig `mapstructure:"mongodb"`
}

func LoadConfig() (config Config, err error) {
	// env := os.Getenv("GO_ENV")

	// if env == "production" {
	// 	return Config{
	// 		Port: os.Getenv("PORT"),
	// 	}, nil
	// }

	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("toml")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	// viper.Debug()
	err = viper.Unmarshal(&config)

	return
}

func (c *Config) Debug() {
	fmt.Printf("%+v\n", c)
}
