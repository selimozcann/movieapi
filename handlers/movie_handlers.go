package handlers

import (
	"log"
	"movie/api/config"
	DatabaseFactory "movie/api/database"
	MovieModel "movie/api/models"

	"github.com/gofiber/fiber/v2"
)

type MovieHandlers struct {
	Db DatabaseFactory.MovieOperations
}

var MovieH *MovieHandlers = &MovieHandlers{}

func init() {
	movieOperation, err := DatabaseFactory.MovieOperationsFactory(config.DB_TYPE)
	if err != nil {
		log.Fatal(err)
	}
	MovieH.Db = movieOperation
}

func (movieH *MovieHandlers) GetMovie(c *fiber.Ctx) error {
	movieID := c.Query("_id")
	if movieID == "" {
		return c.Status(500).JSON(fiber.Map{"status": "unsuccessful", "is_success": false, "message": "Movie ID must be provided."})
	}
	mModel, err := movieH.Db.GetMovie(movieID, c)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "unsuccessful", "is_success": false, "message": err.Error()})
	}
	return c.JSON(mModel)
}
func (movieH *MovieHandlers) GetAllMovies(c *fiber.Ctx) error {
	mModel, err := movieH.Db.GetAllMovies(c)
	if err != nil {
		log.Fatal(err.Error())
	}
	return c.Status(500).JSON(fiber.Map{"get movie is succesfully": mModel})
}
func UpdateMovie() {
	// TO DO Update Movie
}
func (movieH *MovieHandlers) DeleteMovies(c *fiber.Ctx) error {
	movieID := c.Query("_id")
	if movieID == "" {
		return c.Status(500).JSON(fiber.Map{"status": "unsuccessful", "is_success": false, "message": "Movie ID must be provided."})
	}
	err := movieH.Db.DeleteMovie(movieID, c)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "unsuccessful", "is_success": false, "message": err.Error()})
	}
	return c.JSON(fiber.Map{"status": "successful", "message": "Successfully deleted post.", "is_success": true})
}
func (movieH *MovieHandlers) CreateMovies(c *fiber.Ctx) error {
	movieModel := &MovieModel.Movie{}
	if err := c.BodyParser(movieModel); err != nil {
		return c.Status(400).SendString("Invalid request body")
	}
	movieCreated, err := movieH.Db.CreateMovie(*movieModel, c)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(movieCreated)
}
