package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vishalpandhare01/school_management_system/db"
	"github.com/vishalpandhare01/school_management_system/internals/model"
)

type Body struct {
	ID   string
	Name string
	Use  string
}

func AddData(c *fiber.Ctx) error {
	var body Body

	if err := c.BodyParser(&body); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	var data = model.Programing{
		ID:   body.ID,
		Name: body.Name,
		Use:  body.Use,
	}

	dataBase := db.DatabaseConnection()
	if err := dataBase.Create(&data).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Hello world",
		"data":    data,
	})

}

func GetData(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"message": "Hello world",
	})

}
