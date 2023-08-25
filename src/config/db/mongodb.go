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

var (
	mongoOnce sync.Once
	mongoDB   DbHandler
	mongoErr  error
)

const TIMEOUT = 10

func connect(ctx context.Context) (DbHandler, error) {
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

func GetConnection() (DbHandler, context.Context, error) {
	ctx, _ := context.WithTimeout(context.Background(), TIMEOUT*time.Second)
	// defer cancel()

	var getMongoConnection = func() {
		connect(ctx)
	}

	mongoOnce.Do(getMongoConnection)

	connectionErr, db := mongoErr, mongoDB
	fmt.Println(connectionErr, db)

	if connectionErr != nil {
		return nil, nil, connectionErr
	}

	return db, ctx, nil
}
