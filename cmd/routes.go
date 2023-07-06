package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vaibhavvvvv/go-trivia-api/handlers"
)

func setupRoutes(app *fiber.App) {

	app.Get("/", handlers.ListFacts)
	app.Post("/fact", handlers.CreateFact)
	app.Delete("/delete/:id", handlers.DeleteFact)
	app.Put("/fact/:id", handlers.UpdateFact)

}
