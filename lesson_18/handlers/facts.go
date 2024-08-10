package handlers

import (
	"docker_postgres/database"
	"docker_postgres/models"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func ListFacts(c *fiber.Ctx) error {
	db := database.GetDB()
	facts := []models.Fact{}
	ff := db.Find(&facts)
	if ff.Error != nil {
		return fmt.Errorf("error during get facts")
	}
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

	f := db.Create(&fact)
	if f.Error != nil {
		return fmt.Errorf("error during saving fact with id: %v ", fact)
	}

	return c.Status(200).JSON(fact)
}
