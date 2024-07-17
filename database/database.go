package database

import (
	"fmt"

	"github.com/gporawetku/go-blog-api-example/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		config.AppConfig.DBHost, config.AppConfig.DBPort, config.AppConfig.DBUser, config.AppConfig.DBName, config.AppConfig.DBPassword)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
