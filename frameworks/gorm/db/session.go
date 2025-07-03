package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDB() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=admin dbname=gorm port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}
