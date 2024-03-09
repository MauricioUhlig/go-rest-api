package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectToDB() {
	connStr := "user=go_user dbname=go_learning password=123456789 host=192.168.100.100 port=32769 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Panic("Error to connect to database!")
	}
}
