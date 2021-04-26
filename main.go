package main

import (
	"fmt"
	"os"

	"github.com/aunza007/env"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	switch os.Getenv("ENV") {
	case "pro":
		fmt.Println("Pro")
	default:
		fmt.Println("Dev")
		env.SetDefaultEnv()
	}

	app := fiber.New(fiber.Config{
		BodyLimit:     25 * 1024 * 1024,
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
	})
	app.Use(cors.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))
	app.Use(logger.New(logger.Config{
		Format:     "${blue}${time} ${yellow}${status} - ${red}${latency} ${cyan}${method} ${path} ${green} ${ip} ${ua} ${reset}\n",
		TimeFormat: "02-Jan-2006 15:04:05",
		TimeZone:   "Asia/Bangkok",
		Output:     os.Stdout,
	}))

	app.Get("/", func(c *fiber.Ctx) error {

		return c.SendString("NGS Chat API v1 " + os.Getenv("ENV"))
	})

	//routes.SetUpRoutes(app)

	app.Listen(":7000")

}
