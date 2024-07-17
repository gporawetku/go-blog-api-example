package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

var AppConfig Config

func LoadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	AppConfig = Config{
		DBHost:     viper.GetString("database.host"),
		DBPort:     viper.GetString("database.port"),
		DBUser:     viper.GetString("database.user"),
		DBPassword: viper.GetString("database.password"),
		DBName:     viper.GetString("database.name"),
	}
}
