package main

import (
    "filebed/routers"

    "github.com/gofiber/fiber/v2"
)

func main() {
	prefork := true
    app := fiber.New(fiber.Config{
    	Prefork: prefork,
    	ServerHeader: "Go",
    	StrictRouting: true,
    	CaseSensitive: true,
    })

    routers.Register(app)

    app.Listen(":3750")
}
