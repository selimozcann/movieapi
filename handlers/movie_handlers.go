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

var movieModel []MovieModel.Movie

func (movieH *MovieHandlers) GetMovie(c *fiber.Ctx) error {
	/*
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
	*/
	return c.SendString("Movie ID not found")
}
func (movieH *MovieHandlers) GetAllMovies(c *fiber.Ctx) error {
	mModel, err := movieH.Db.GetAllMovies(c)
	if err != nil {
		log.Fatal(err.Error())
	}
	return c.Status(500).JSON(fiber.Map{"get movie is succesfully": mModel})
}
func updateMovie() {
	// Update Init UpdateMovie
}
func (movieH *MovieHandlers) DeleteMovies(c *fiber.Ctx) error {
	/*
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
	*/
	return c.SendString("Movie not found")
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
