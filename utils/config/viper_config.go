package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	ServerHost string
	ServerPort string
}

func InitConfig() (*Config, error) {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	config := setConfig()
	return config, nil
}

func setConfig() *Config {
	config := new(Config)
	config.ServerHost = viper.GetString("SERVER_HOST")
	config.ServerPort = viper.GetString("SERVER_PORT")
	return config
}
