package db

import (
	"log"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"migrations/models"
)

func Migrate() {
	dsn := "host=localhost user=postgres password=admin dbname=gorm port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "20250703_create_users",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&models.User{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("users")
			},
		},
		{
			ID: "20250704_create_products",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&models.Product{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("products")
			},
		},
	})

	if err := m.Migrate(); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	log.Println("Migration applied")
}
