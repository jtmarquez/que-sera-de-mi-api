package mongodb_handler

import (
	"context"
	"fmt"
	dotenv "que-sera-de-mi/src/config/env_vars"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DbHandler = *mongo.Database
type DbCommandCallback func(ctx context.Context, db DbHandler) error

var (
	mongoOnce sync.Once
	mongoDB   DbHandler
	mongoErr  error
)

const TIMEOUT = 10

// TODO: Check connection is created and still valid, otherwise reconnect.
func getConnection(ctx context.Context) (DbHandler, error) {
	dbUrl := dotenv.GetEnvironmentVariable("DB_URL")
	dbName := dotenv.GetEnvironmentVariable("DB_NAME")

	options := options.Client().ApplyURI(dbUrl)

	client, err := mongo.Connect(ctx, options)

	if err != nil {
		mongoErr = err
		return nil, err
	}

	mongoDB = client.Database(dbName)
	fmt.Println(mongoDB)
	return mongoDB, nil
}

func ExecuteDbCommand(callback DbCommandCallback) error {
	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT*time.Second)
	defer cancel()

	var getMongoConnection = func() {
		getConnection(ctx)
	}

	fmt.Println(mongoDB)
	mongoOnce.Do(getMongoConnection)
	fmt.Println(mongoDB)

	connectionErr, db := mongoErr, mongoDB

	if connectionErr != nil {
		return connectionErr
	}

	err := callback(ctx, db)

	if err != nil {
		return err
	}

	defer db.Client().Disconnect(ctx)

	return nil
}
