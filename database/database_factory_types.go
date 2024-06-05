package database

import (
	MovieModel "movie/api/models"

	"github.com/gofiber/fiber/v2"
)

type MovieOperations interface {
	GetMovie(movieID string, c *fiber.Ctx) (MovieModel.Movie, error)
	GetAllMovies(c *fiber.Ctx) ([]MovieModel.Movie, error)
	CreateMovie(movieModel MovieModel.Movie, c *fiber.Ctx) (MovieModel.Movie, error)
	UpdateMovie(movieID string, c *fiber.Ctx) error
	DeleteMovie(movieID string, c *fiber.Ctx) error
}
