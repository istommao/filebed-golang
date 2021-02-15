package apis

import (
    "fmt"
    "github.com/gofiber/fiber/v2"
)


func UploadAPI(c *fiber.Ctx) error {
    file, err := c.FormFile("file")
    if err != nil {
        return err
    }

    c.SaveFile(file, fmt.Sprintf("media/%s", file.Filename))

    return c.SendString("Upload success 👋!")
}
