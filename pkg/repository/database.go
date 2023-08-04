package repository

import (
	"fmt"

	config "stream/pkg/config"
	"stream/pkg/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() (*gorm.DB, error) {
	config, err := config.LoadConfig("./")
	db, err := gorm.Open(postgres.Open(config.DSN), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	DB = db
	DB.AutoMigrate(models.Video{})
	return db, nil
}
