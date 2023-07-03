package database

import (
	"errors"
	MovieOperationsMongo "movie/api/database/mongooperations/post"
)

const (
	mongoDb = "mongo"
)

func MovieOperationFactory(dbType string) (MovieOperations, error) {
	if dbType == mongoDb {
		return MovieOperationsMongo.MovieOperationObj, nil
	}
	return nil, errors.New("DB Type not found")
}
