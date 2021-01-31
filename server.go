package main

import (
    "filebed/routers"
    "filebed/config"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/template/html"
)

func main() {
    conf := config.InitConfigure()

    engine := html.New("./templates", ".html")

    app := fiber.New(fiber.Config{
    	Prefork: conf.Prefork,
    	ServerHeader: "Go",
    	StrictRouting: true,
    	CaseSensitive: true,
        Views: engine,
    })

    routers.Register(app)

    app.Listen(conf.Host + ":" + conf.Port)
}
