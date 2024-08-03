package main

import (
	"docker_postgres/database"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	if err := database.Connect(); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	if err := database.AutoMigrate(); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	setupRoutes(app)

	app.Listen(":3000")
}
