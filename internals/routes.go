package internals

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/vishalpandhare01/school_management_system/internals/handler"
)

func SetUpRoutes(app *fiber.App) {
	app.Get("/", handler.GetData)
	app.Post("/", handler.AddData)

	fmt.Println("Routes Connected successfully!!!")
}
