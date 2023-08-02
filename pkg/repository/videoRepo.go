package repository

import (
	"errors"
	"log"
	"stream/pkg/models"

	"gorm.io/gorm"
)

type VideoRepo interface {
}

func init() {
	DB, err = ConnectDatabase()
	if err != nil {
		log.Fatal(err)
	}
}

func Create(input *models.Video) error {
	return DB.Create(input).Error
}

func FindBySectionID(input int) (string, error) {
	var video models.Video
	result := DB.Where("section =?", input).First(&video)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return "", errors.New("cart not found")
		}
		return "", errors.New("cart not found")
	}
	return video.FileName, nil
}
