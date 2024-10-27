package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/vishalpandhare01/school_management_system/db"
	"github.com/vishalpandhare01/school_management_system/internals"
	"github.com/vishalpandhare01/school_management_system/migration"
)

func main() {
	app := fiber.New()
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	db.DatabaseConnection()
	err := migration.Migration()
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:8000, http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept , Token , Authorization",
	}))

	internals.SetUpRoutes(app)

	app.Listen(":8000")
}
