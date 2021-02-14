package apis

import (
	"fmt"
    "github.com/gofiber/fiber/v2"
)

func IndexPage(c *fiber.Ctx) error {
    return c.Render("index", fiber.Map{
        "title": "Hello, World!",
    })
}


func HomePage(c *fiber.Ctx) error {
    return c.Render("home", fiber.Map{
        "title": "home",
    })
}

func MediaPage(c *fiber.Ctx) error {
    return c.Render("media", fiber.Map{
        "title": "Media",
    })
}


func UploadAPI(c *fiber.Ctx) error {
    file, err := c.FormFile("file")
    if err != nil {
        return err
    }

    c.SaveFile(file, fmt.Sprintf("media/%s", file.Filename))

    return c.SendString("Upload success 👋!")
}
