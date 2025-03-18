package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/orbulant/stock-screener/internal/routes"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	routes.SetupRoutes(app)

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("PORT not specified")
	}

	app.Listen(":" + port)
}
