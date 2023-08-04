package video

import (
	"errors"
	"log"

	"fmt"

	"gorm.io/gorm"

	config "stream/pkg/config"

	"gorm.io/driver/postgres"
)

var DB *gorm.DB
var err error

func ConnectDatabase() (*gorm.DB, error) {
	config, err := config.LoadConfig("./")
	db, err := gorm.Open(postgres.Open(config.DSN), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	DB = db
	DB.AutoMigrate(Video{})
	return db, nil
}

type VideoRepo interface {
}

func init() {
	DB, err = ConnectDatabase()
	if err != nil {
		log.Fatal(err)
	}
}

func Create(input *Video) error {
	return DB.Create(input).Error
}

func FindBySectionID(input int) (string, error) {
	var video Video
	result := DB.Where("section =?", input).First(&video)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return "", errors.New("cart not found")
		}
		return "", errors.New("cart not found")
	}
	return video.FileName, nil
}
