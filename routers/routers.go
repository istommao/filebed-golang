package routers

import (
	"filebed/apis"
	"github.com/gofiber/fiber/v2"
)


func Register(app *fiber.App) {
    app.Get("/", apis.IndexPage)
    app.Get("/upload/", apis.UploadPage)
    app.Get("/files/", apis.MediaPage)

    app.Static("/static", "./static")

	api := app.Group("/api")
    api.Post("/upload/", apis.UploadAPI)
}
