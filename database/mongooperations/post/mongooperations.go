package post

import (
	"errors"
	"log"
	Mongo "movie/api/database/mongo"
	MovieModel "movie/api/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MovieOperationsMongo struct {
	db *Mongo.MongoInstance
}

var MovieOperationObj *MovieOperationsMongo

func init() {
	mongoDb := Mongo.GetDatabase()
	MovieOperationObj = &MovieOperationsMongo{db: mongoDb}
}

func (movieOper *MovieOperationsMongo) GetAllMovies(c *fiber.Ctx) ([]MovieModel.Movie, error) {
	movies := make([]MovieModel.Movie, 0)

	query := bson.D{{}}
	cursor, err := movieOper.db.Db.Collection("movie").Find(c.Context(), query)
	if err != nil {
		return movies, err
	}
	if err := cursor.All(c.Context(), &movies); err != nil {
		return movies, err
	}
	return movies, nil
}
func (movieOper *MovieOperationsMongo) GetMovie(movieID string, c *fiber.Ctx) (MovieModel.Movie, error) {
	primitiveObjID, err := primitive.ObjectIDFromHex(movieID)
	if err != nil {
		return MovieModel.Movie{}, err
	}
	filter := bson.D{{Key: "_id", Value: primitiveObjID}}
	result := movieOper.db.Db.Collection("movie").FindOne(c.Context(), filter)

	fetchedMovie := &MovieModel.Movie{}
	err = result.Decode(fetchedMovie)
	if err != nil {
		return MovieModel.Movie{}, err
	}
	return *fetchedMovie, nil
}
func (movieOper *MovieOperationsMongo) CreateMovie(movieModel MovieModel.Movie, c *fiber.Ctx) (MovieModel.Movie, error) {
	createPost, err := movieOper.db.Db.Collection("movie").InsertOne(c.Context(), movieModel)
	if err != nil {
		return MovieModel.Movie{}, err
	}
	filter := bson.D{{Key: "_id", Value: createPost.InsertedID}}
	createRecord := movieOper.db.Db.Collection("movie").FindOne(c.Context(), filter)
	createdPost := &MovieModel.Movie{}

	err = createRecord.Decode(createdPost)
	if err != nil {
		return MovieModel.Movie{}, err
	}
	return *createdPost, nil
}
func (movieOper *MovieOperationsMongo) UpdateMovie(movieId string, c *fiber.Ctx) error {

	movieModel := new(MovieModel.Movie)
	if err := c.BodyParser(movieModel); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}
	primitiveObjID, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid movie ID"})
	}
	query := bson.D{{Key: "_id", Value: primitiveObjID}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "moviename", Value: movieModel.MovieName},
			{Key: "releaseyear", Value: movieModel.ReleaseYear},
			{Key: "genre", Value: movieModel.Genre},
			{Key: "directedby", Value: movieModel.DirectedBy},
		}},
	}
	opts := options.Update()

	r, err := movieOper.db.Db.Collection("movie").UpdateOne(c.Context(), query, update, opts)
	if err != nil {
		log.Println("Error while updating movie: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update movie"})
	}
	log.Println("Successfully updated movie: ", r)
	return c.JSON(r)
}
func (movieOper *MovieOperationsMongo) DeleteMovie(movieID string, c *fiber.Ctx) error {
	primitiveObjID, err := primitive.ObjectIDFromHex(movieID)
	if err != nil {
		return err
	}

	filter := bson.D{{Key: "_id", Value: primitiveObjID}}
	result, err := movieOper.db.Db.Collection("movie").DeleteOne(c.Context(), &filter)
	if err != nil {
		return err
	}
	if result.DeletedCount < 1 {
		return errors.New("Delete count is empty")
	}
	return nil
}
