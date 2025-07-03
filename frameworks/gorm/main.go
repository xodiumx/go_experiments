package main

import (
	"log"
	"migrations/db"
	"migrations/models"
	"time"
)

func main() {
	db.Migrate()

	gormDB, err := db.SetupDB()
	if err != nil {
		log.Fatal(err)
	}

	// Настроим пул соединений
	sqlDB, err := gormDB.DB()
	if err != nil {
		log.Fatal(err)
	}

	sqlDB.SetMaxIdleConns(10)           // сколько "пассивных" соединений держать
	sqlDB.SetMaxOpenConns(100)          // максимальное количество одновременных соединений
	sqlDB.SetConnMaxLifetime(time.Hour) // сколько живёт соединение

	// CREATE
	user := models.User{Name: "Maksim", Email: "maks7@example.com"}
	result := gormDB.Create(&user)
	if result.Error != nil {
		log.Fatal(result.Error)
	} else {
		log.Printf("User created successfully: %v", user)
	}

	// GET
	var newUser models.User
	gormDB.First(&newUser, 1) // SELECT * FROM users WHERE id = 1 LIMIT 1;
	gormDB.First(&newUser, "email = ?", "maks@example.com")
	log.Printf("User info: %v", newUser)

	// UPDATE
	//gormDB.Model(&user).Update("name", "Max Updated")
	// Many:
	//gormDB.Model(&user).Updates(models.User{Name: "NewName", Email: "new2@example.com"})

	// DELETE
	gormDB.Delete(&user)
	// или по ID
	//gormDB.Delete(&models.User{}, 1)

}
