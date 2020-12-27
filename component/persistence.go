package component

import (
	"fmt"
	"github.com/JIeeiroSst/store/models/abouts"
	"github.com/JIeeiroSst/store/models/categories"
	"github.com/JIeeiroSst/store/models/contacts"
	feedBacks "github.com/JIeeiroSst/store/models/feed-backs"
	"github.com/JIeeiroSst/store/models/menues"
	"github.com/JIeeiroSst/store/models/news"
	"github.com/JIeeiroSst/store/models/product_category"
	"github.com/JIeeiroSst/store/models/products"
	"github.com/JIeeiroSst/store/models/sliders"
	"github.com/JIeeiroSst/store/models/system_config"
	"github.com/JIeeiroSst/store/models/tags"
	"github.com/JIeeiroSst/store/models/users"
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
	conn.Debug().AutoMigrate(&abouts.Abouts{},&categories.Categories{},&contacts.Contacts{},&feedBacks.FeedBacks{},&menues.Menues{},&news.News{},&product_category.ProductCategory{},&products.Products{},&sliders.Sliders{},&system_config.SystemConfig{},&tags.Tags{},&users.Users{})

	GlobalCache, err = bigcache.NewBigCache(bigcache.DefaultConfig(30 * time.Minute))
	if err != nil {
		panic(fmt.Errorf("failed to initialize cahce: %w", err))
	}
}

func GetConn() *gorm.DB {
	return DB
}
