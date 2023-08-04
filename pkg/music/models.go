package music

import "gorm.io/gorm"

type Music struct {
	gorm.Model  `json:"-"`
	Section     int    `json:"section"`
	FileName    string `json:"filename"`
	Url         string `json:"url"`
	Description string `json:"description"`
}
