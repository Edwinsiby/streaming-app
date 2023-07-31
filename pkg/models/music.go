package models

type Music struct {
	FileName    string `json:"filename"`
	Url         string `json:"url"`
	Description string `json:"description"`
}
