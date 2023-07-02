package handlers

import (
	"log"
	MovieModel "movie/api/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

var movieModel []MovieModel.Movie

func getMovie(c *fiber.Ctx) error {
	p := c.Params("id")
	if p == "" {
		c.SendString("Id is empty")
	}
	intP, _ := strconv.Atoi(p)
	for _, mModel := range movieModel {
		if mModel.Id == intP {
			return c.JSON(mModel)
		}
	}
	return c.SendString("Movie ID not found")
}
func getAllMovies(c *fiber.Ctx) error {
	return c.JSON(movieModel)
}
func updateMovie() {
	// Update Init UpdateMovie
}
func deleteMovies(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.SendString("Id is empty")
	}
	intP, _ := strconv.Atoi(id)
	for i, mModel := range movieModel {
		if mModel.Id == intP {
			movieModel = append(movieModel[:i], movieModel[i+1:]...)
			return c.SendString("Movie deleted")
		}
	}
	return c.SendString("Movie not found")
}
func createMovies(c *fiber.Ctx) error {
	mModel := &MovieModel.Movie{}
	err := c.BodyParser(mModel)
	if err != nil {
		c.SendString(err.Error())
	}
	movieModel = append(movieModel, *mModel)
	return c.JSON(movieModel)
}
func InitEndpoint() {
	app := fiber.New()
	app.Get("/getMovies", getAllMovies)
	app.Get("/getMovie", getMovie)
	app.Get("/getMovie/:id", getMovie)

	app.Post("/createMovies", createMovies)

	app.Delete("/delete", deleteMovies)
	app.Delete("/delete/:id", deleteMovies)

	log.Fatal(app.Listen(":8000"))
}
