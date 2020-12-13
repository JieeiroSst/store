package component

import (
	"fmt"
	"github.com/allegro/bigcache"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

var (
	DB          *gorm.DB
	GlobalCache *bigcache.BigCache
)

func init() {
	// Connect to DB
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}
	username := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")
	dbUri := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, username, dbName, password)
	conn,err:=gorm.Open(postgres.Open(dbUri),&gorm.Config{})
	if err != nil {
		log.Println("failed to connect database")
	}
	log.Println("server connect database success")
	DB = conn
	//db.Debug().AutoMigrate(&abouts.Abouts{},&categories.Categories{},&contacts.Contacts{},&feed_backs.FeedBacks{},&menues.Menues{},&news.News{},&news.News{},&product_category.ProductCategory{},&products.Products{},&sliders.Sliders{},&system_config.SystemConfig{},&tags.Tags{},&users.Users{})

	// Initialize cache
	GlobalCache, err = bigcache.NewBigCache(bigcache.DefaultConfig(30 * time.Minute))
	if err != nil {
		panic(fmt.Errorf("failed to initialize cahce: %w", err))
	}
}

func GetConn() *gorm.DB {
	return DB
}
