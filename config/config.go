package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Api   Api
	Mongo Mongo
	Redis Redis
}
type Api struct {
	Port int64
	Url  string
}
type Mongo struct {
	Host     string
	Port     int32
	User     string
	Password string
}
type Redis struct {
	Host     string
	Port     int32
	User     string
	Password string
	Database int
}

var config Config

func LoadConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("fatal error config file: %w", err)
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		return err
	}

	return nil
}

func GetConfig() Config {
	return config
}
