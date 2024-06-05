package post

import (
	"errors"
	"log"
	Mongo "movie/api/database/mongo"
	MovieModel "movie/api/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
func (movieOper *MovieOperationsMongo) UpdateMovie(movieID string, c *fiber.Ctx) error {
	primitiveObjID, err := primitive.ObjectIDFromHex(movieID)
	if err != nil {
		return err
	}

	update := bson.D{
		{"$set", bson.D{
			{"moviename", c.FormValue("moviename")},
			{"releaseyear", c.FormValue("releaseyear")},
			{"genre", c.FormValue("genre")},
			{"directedby", c.FormValue("directedby")},
		}},
	}

	r, err := movieOper.db.Db.Collection("movie").UpdateByID(c.Context(), primitiveObjID, update)
	if err != nil {
		log.Println("Error while updating movie: ", err)
		return err
	}
	log.Println("Succesfully updated movie: ")
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
