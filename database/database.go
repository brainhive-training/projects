package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const DSN = "host=cluster-example-rw user=project_api password=b00kworM dbname=project_api port=5432 sslmode=disable TimeZone=Europe/Amsterdam"

func Connect() *gorm.DB {
	conn, err := gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return conn
}
