package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	HTTP struct {
		Host string
		Port int
	}
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config") // config.yaml
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var c Config
	if err := viper.Unmarshal(&c); err != nil {
		return nil, err
	}
	return &c, nil
}
