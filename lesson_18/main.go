package main

import (
	"docker_postgres/database"
	"flag"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

// func main() {
// 	app := fiber.New()
// 	if err := database.Connect(); err != nil {
// 		log.Fatal("Failed to connect to database:", err)
// 	}

// 	if err := database.AutoMigrate(); err != nil {
// 		log.Fatal("Failed to migrate database:", err)
// 	}

// 	setupRoutes(app)

// 	app.Listen(":3000")
// }

func main() {
	migrateCmd := flag.NewFlagSet("migrate", flag.ExitOnError)
	serveCmd := flag.NewFlagSet("serve", flag.ExitOnError)

	if len(os.Args) < 2 {
		log.Fatal("expected 'migrate' or 'serve' subcommands")
	}

	switch os.Args[1] {
	case "migrate":
		migrateCmd.Parse(os.Args[2:])
		runMigrate()
	case "serve":
		serveCmd.Parse(os.Args[2:])
		runServer()
	default:
		log.Fatal("expected 'migrate' or 'serve' subcommands")
	}
}

func runMigrate() {
	if err := database.Connect(); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	if err := database.AutoMigrate(); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Database migrated successfully")
}

func runServer() {
	if err := database.Connect(); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	app := fiber.New()

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
