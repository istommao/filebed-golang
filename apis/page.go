package apis

import (
    "github.com/gofiber/fiber/v2"
)

func IndexPage(c *fiber.Ctx) error {
    return c.Render("index", fiber.Map{
        "title": "Hello, World!",
    })
}

func UploadPage(c *fiber.Ctx) error {
    return c.Render("upload", fiber.Map{
        "title": "upload",
    })
}

func MediaPage(c *fiber.Ctx) error {
    return c.Render("media", fiber.Map{
        "title": "Media",
    })
}
