package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const DSN = "host=localhost user=project_api password=b00kworM dbname=project_api port=9000 sslmode=disable TimeZone=Europe/Amsterdam"

func Connect() *gorm.DB {
	conn, err := gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return conn
}
