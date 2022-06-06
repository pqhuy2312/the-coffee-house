package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/pqhuy2312/the-coffee-house/configs"
	"github.com/pqhuy2312/the-coffee-house/routers"
)

func main() {
	err := godotenv.Load(".env")
    if err != nil {
        log.Println("Load .env error")
    }
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	app := fiber.New()



	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowCredentials: true,
	}))

	routers.InitializeApiMapping(app)

	configs.AutoMigrate()


	log.Fatal(app.Listen(":"+port))
}