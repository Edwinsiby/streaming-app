package video

import "gorm.io/gorm"

type Video struct {
	gorm.Model  `json:"-"`
	Section     int    `json:"section"`
	FileName    string `json:"filename"`
	Url         string `json:"url"`
	Description string `json:"description"`
}
