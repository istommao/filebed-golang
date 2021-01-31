package apis

import (
	"fmt"
    "github.com/gofiber/fiber/v2"
)


func HomePage(c *fiber.Ctx) error {
    return c.SendString("Hello, World 👋!")
}

func UploadAPI(c *fiber.Ctx) error {
    file, err := c.FormFile("document")
    if err != nil {
        return err
    }

    c.SaveFile(file, fmt.Sprintf("static/%s", file.Filename))

    return c.SendString("Upload success 👋!")
}
