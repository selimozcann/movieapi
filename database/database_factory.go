package database

import (
	"errors"
	MovieOperationsMongo "movie/api/database/mongooperations/post"
)

const (
	mongoDB = "mongo"
)

func MovieOperationsFactory(dbType string) (MovieOperations, error) {
	if dbType == mongoDB {
		return MovieOperationsMongo.MovieOperationObj, nil
	}
	return nil, errors.New("DB Type is not found")
}
