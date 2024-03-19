package route

import (
	"github.com/gofiber/fiber/v2"

	"github.com/iuan95/apigo/handler"
)


func InitRoute(app *fiber.App){

	app.Get("/", handler.GetItems)
	app.Get("/:id", handler.GetItemByID)
}