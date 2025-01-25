package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Database struct {
		Host     string `mapstructure:"DB_HOST"`
		Port     string `mapstructure:"DB_PORT"`
		User     string `mapstructure:"DB_USER"`
		Password string `mapstructure:"DB_PASSWORD"`
		Name     string `mapstructure:"DB_NAME"`
	} `mapstructure:"database"`
	Redis struct {
		Host     string `mapstructure:"REDIS_HOST"`
		Port     string `mapstructure:"REDIS_PORT"`
		Password string `mapstructure:"REDIS_PASSWORD"`
	} `mapstructure:"redis"`
}

func LoadConfig() (*Config, error) {
	var config Config

	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
		return nil, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
		return nil, err
	}

	return &config, nil
}