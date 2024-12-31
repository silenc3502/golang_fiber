package main

import (
	"fmt"
	"golang_fiber/initializer"
	"golang_fiber/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// DB 초기화
	db, err := initializer.DomainInitializer()
	if err != nil {
		fmt.Println("Error initializing domain:", err)
		return
	}

	// Initialize the Fiber app
	app := fiber.New()

	// Register routes for all domains (posts, users, etc.)
	router.RegisterRoutes(app, db)

	// Start the server
	if err := app.Listen(":3773"); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
