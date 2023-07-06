package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vaibhavvvvv/go-trivia-api/database"
	"github.com/vaibhavvvvv/go-trivia-api/models"
)

func ListFacts(c *fiber.Ctx) error {
	facts := []models.Fact{}
	database.DB.Db.Find(&facts)

	return c.Status(200).JSON(facts)

}

func CreateFact(c *fiber.Ctx) error {
	fact := new(models.Fact)
	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&fact) //persist the fact(data) to our database

	return c.Status(200).JSON(fact)
}

func DeleteFact(c *fiber.Ctx) error {
	id := c.Params("id")
	fact := models.Fact{}
	database.DB.Db.Where("id = ?", id).First(&fact)

	if fact.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Fact not found",
		})
	}

	database.DB.Db.Delete(&fact)

	return c.Status(200).JSON(fiber.Map{
		"message": "Fact deleted",
	})
}

func UpdateFact(c *fiber.Ctx) error {
	id := c.Params("id")
	fact := models.Fact{}

	database.DB.Db.Where("id = ?", id).First(&fact)

	if fact.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Fact not found",
		})
	}

	if err := c.BodyParser(&fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Updates(&fact)

	return c.Status(200).JSON(fact)

}
