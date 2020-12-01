package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jeeiro/api-go/models/abouts"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var (
	db  *gorm.DB
)


func init() {
	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")

	dbUri := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, username, dbName, password)

	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		log.Println("failed to connect database")
	}

	log.Println("server connect database success")

	db = conn
	// db.Debug().AutoMigrate(&models.Account{}, &models.Role{}, &models.Profile{})
}

func ConnDB() *gorm.DB {
	return db
}