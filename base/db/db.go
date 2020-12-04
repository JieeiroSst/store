package db

import (
	"fmt"
	"github.com/jieeiro/api/models/abouts"
	"github.com/jieeiro/api/models/categories"
	"github.com/jieeiro/api/models/contacts"
	"github.com/jieeiro/api/models/feed_backs"
	"github.com/jieeiro/api/models/menues"
	"github.com/jieeiro/api/models/news"
	"github.com/jieeiro/api/models/product_category"
	"github.com/jieeiro/api/models/products"
	"github.com/jieeiro/api/models/sliders"
	"github.com/jieeiro/api/models/system_config"
	"github.com/jieeiro/api/models/tags"
	"github.com/jieeiro/api/models/users"
	"log"
	"os"

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
	db.Debug().AutoMigrate(&abouts.Abouts{},&categories.Categories{},&contacts.Contacts{},&feed_backs.FeedBacks{},&menues.Menues{},&news.News{},&news.News{},&product_category.ProductCategory{},&products.Products{},&sliders.Sliders{},&system_config.SystemConfig{},&tags.Tags{},&users.Users{})
}

func GetConn() *gorm.DB {
	return db
}