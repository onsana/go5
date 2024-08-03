package handlers

import (
	"docker_postgres/database"
	"docker_postgres/models"

	"github.com/gofiber/fiber/v2"
)

func ListFacts(c *fiber.Ctx) error {
	db := database.GetDB()
	facts := []models.Fact{}
	db.Find(&facts)

	return c.Status(200).JSON(facts)
}

func CreateFact(c *fiber.Ctx) error {
	db := database.GetDB()
	fact := new(models.Fact)
	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	db.Create(&fact)

	return c.Status(200).JSON(fact)
}
