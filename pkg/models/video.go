package models

type Video struct {
	FileName    string `json:"filename"`
	Url         string `json:"url"`
	Description string `json:"description"`
}
