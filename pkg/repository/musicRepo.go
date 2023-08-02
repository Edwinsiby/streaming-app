package repository

import (
	"log"
)

type MusicRepo interface {
}

var err error

func init() {
	DB, err = ConnectDatabase()
	if err != nil {
		log.Fatal(err)
	}
}
