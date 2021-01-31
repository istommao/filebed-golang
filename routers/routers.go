package routers

import (
	"filebed/apis"
	"github.com/gofiber/fiber/v2"
)


func Register(app *fiber.App) {
    app.Get("/", apis.HomePage)

    app.Static("/static", "./static")

	api := app.Group("/api")
    api.Post("/upload/", apis.UploadAPI)
}
