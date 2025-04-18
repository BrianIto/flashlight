package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func InitDatabase() *gorm.DB {
	dsn := "host=localhost user=myuser password=mypassword dbname=mydatabase port=5432 sslmode=disable TimeZone=UTC"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	log.Println("Connected to PostgreSQL with GORM!")

	return db
}
