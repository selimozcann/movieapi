package main

import (
	"log"
	"movie/api/config"
	"movie/api/routers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	routers.AddMovieRouters(app)

	err := app.Listen(config.PORT)
	if err != nil {
		log.Fatal(err)
	}
}
