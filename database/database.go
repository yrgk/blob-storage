package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDb() *gorm.DB {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Print("No .env file found")
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASS")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%s", host, user, dbname, password, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}