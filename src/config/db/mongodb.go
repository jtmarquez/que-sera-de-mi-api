package mongodb_handler

import (
	"context"
	dotenv "que-sera-de-mi/src/config/env_vars"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DbHandler = *mongo.Database
type DbCommandCallback func(ctx context.Context, db DbHandler) (interface{}, error)

var (
	mongoOnce sync.Once
	mongoDB   DbHandler
	mongoErr  error
)

const TIMEOUT = 10

func getConnection(ctx context.Context) (DbHandler, error) {
	dbUrl := dotenv.GetEnvironmentVariable("DB_URL")
	dbName := dotenv.GetEnvironmentVariable("DB_NAME")

	options := options.Client().ApplyURI(dbUrl)

	client, err := mongo.Connect(ctx, options)

	if err != nil {
		mongoErr = err
		return nil, err
	}

	if pingErr := client.Ping(ctx, nil); pingErr != nil {
		return nil, pingErr
	}

	mongoDB = client.Database(dbName)

	return mongoDB, nil
}

func ExecuteDbCommand(callback DbCommandCallback) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT*time.Second)
	defer cancel()

	var getMongoConnection = func() {
		getConnection(ctx)
	}

	mongoOnce.Do(getMongoConnection)

	connectionErr, db := mongoErr, mongoDB

	if connectionErr != nil {
		return nil, connectionErr
	}

	result, err := callback(ctx, db)

	if err != nil {
		return nil, err
	}

	return result, nil
}
