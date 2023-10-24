package mongo

import (
	"context"
	"fmt"
	"movie/api/config"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var once sync.Once
var singleInstance *MongoInstance = &MongoInstance{Client: nil, Db: nil}
var mongoURI string

func Connect(mg *MongoInstance) error {
	mongoURI := fmt.Sprintf("%s%s:%s%s", config.MONGO_URL_START, config.MONGO_USER_NAME, config.MONGO_PASSWORD, config.MONGO_DB_URL)

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongoURI).SetServerAPIOptions(serverAPI)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)

	defer cancel()
	client, err := mongo.Connect(ctx, opts)

	if err != nil {
		panic(err)
	}

	db := client.Database(config.MONGO_DB_NAME)

	mg.Client = client
	mg.Db = db
	if clientErr := client.Database("admin").RunCommand(ctx, bson.D{{"ping", 1}}).Err(); clientErr != nil {
		panic(clientErr)
	}

	return nil
}

func GetDatabase() *MongoInstance {
	once.Do(func() {
		err := Connect(singleInstance)
		if err != nil {
			panic(err)
		}
	})
	return singleInstance
}
