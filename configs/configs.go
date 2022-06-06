package configs

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"github.com/pqhuy2312/the-coffee-house/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

func GetConnection() *gorm.DB {
	once.Do(func () {
		envErr := godotenv.Load(".env")
    if envErr != nil {
        log.Fatalln("Load .env error")
    }
	dsn := os.Getenv("DATABASE_URL")
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		

		if err != nil {
			log.Fatalln(err)
		}


	})
	
	return db
}

func AutoMigrate() {
	db := GetConnection()
	db.AutoMigrate(
		models.User{},
		models.Category{},
		models.Product{},
		models.ProductImage{},
		models.Size{},
		models.Topping{},
		models.UserAddress{},
		models.Topic{},
		models.Tag{},
		models.Post{},
	)

}