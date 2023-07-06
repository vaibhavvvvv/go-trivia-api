package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vaibhavvvvv/go-trivia-api/database"
)

// import main() {
// 	app := fiber.New()

// 	app.Get("/", func(c *fiber.Ctx) error {
// 		return c.SendString("Hello, Vaibhav")
// 	} )

//		app.Listen(":3000")
//	}
func main() {
	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)

	app.Listen(":3000")
}
