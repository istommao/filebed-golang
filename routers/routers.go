package routers

import (
	"filebed/apis"
	"github.com/gofiber/fiber/v2"
)


func Register(app *fiber.App) {
    app.Get("/", apis.IndexPage)
    app.Get("/home/", apis.HomePage)
    app.Get("/files/", apis.MediaPage)

    app.Static("/static", "./static")

	api := app.Group("/api")
    api.Post("/upload/", apis.UploadAPI)
}
