package routers

import (
	MovieHandlers "movie/api/handlers"

	"github.com/gofiber/fiber/v2"
)

func AddMovieRouters(app *fiber.App) {
	movie_routees := app.Group("/api/v1/movie")
	// POST
	movie_routees.Post("createMovie", MovieHandlers.MovieH.CreateMovies)

	// DELETE
	movie_routees.Delete("deleteMovie", MovieHandlers.MovieH.DeleteMovies)
	movie_routees.Delete("deleteMovie/:id", MovieHandlers.MovieH.DeleteMovies)

	// GET
	movie_routees.Get("/getMovie", MovieHandlers.MovieH.GetMovie)
	movie_routees.Get("/getMovie/:id", MovieHandlers.MovieH.GetMovie)
	movie_routees.Get("/getAllMovies", MovieHandlers.MovieH.GetAllMovies)
}
