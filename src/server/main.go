package main

import (
	"log"

	"exxo.com/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	db, err := setupAndConnectDB()
	if err != nil {
		panic(err)
	}

	app := fiber.New()

	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET,POST",
	}))
	app.Use(logger.New())

	app.Get("/:shortlink", handlers.Redirect(db))

	v1 := app.Group("/api/v1")

	v1.Get("/shortlinks", handlers.Shortlinks(db))
	v1.Post("/shortlink", handlers.SubmitLink(db))

	log.Fatal(app.Listen(":8800"))
}
