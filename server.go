package main

import (
    "filebed/routers"
    "filebed/config"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/template/html"
	"github.com/markbates/pkger"
)

func main() {
    conf := config.InitConfigure()

	engine := html.NewFileSystem(pkger.Dir("/templates"), ".html")

    fiberConf := fiber.Config {
        Prefork: conf.Prefork,
        ServerHeader: "Go",
        StrictRouting: true,
        CaseSensitive: true,
        Views: engine,
    }

    app := fiber.New(fiberConf)
    
    routers.Register(app)

    app.Listen(conf.Domain)
}
