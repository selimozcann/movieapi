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

	log.Fatal(app.Listen(config.PORT))
}
