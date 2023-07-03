package handlers

import (
	"log"
	"movie/api/config"
	DatabaseFactory "movie/api/database"
	MovieModel "movie/api/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type MovieHandlers struct {
	Db DatabaseFactory.MovieOperations
}

var MovieH *MovieHandlers = &MovieHandlers{}

func init() {
	movieOperation, err := DatabaseFactory.MovieOperationFactory(config.DB_TYPE)
	if err != nil {
		log.Fatal(err)
	}
	MovieH.Db = movieOperation
}

var movieModel []MovieModel.Movie

func (movieH *MovieHandlers) GetMovie(c *fiber.Ctx) error {
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
func (movieH *MovieHandlers) GetAllMovies(c *fiber.Ctx) error {
	return c.JSON(movieModel)
}
func updateMovie() {
	// Update Init UpdateMovie
}
func (movieH *MovieHandlers) DeleteMovies(c *fiber.Ctx) error {
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
func (movieH *MovieHandlers) CreateMovies(c *fiber.Ctx) error {
	mModel := &MovieModel.Movie{}
	err := c.BodyParser(mModel)
	if err != nil {
		c.SendString(err.Error())
	}
	movieModel = append(movieModel, *mModel)
	return c.JSON(movieModel)
}
