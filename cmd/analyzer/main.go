package main

import (
	"os"

	"github.com/samgozman/go-finra-short-sales-analyzer/internal/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	dbname := os.Getenv("MONGODB_NAME")

	credential := options.Credential{
		Username: os.Getenv("MONGO_INITDB_ROOT_USERNAME"),
		Password: os.Getenv("MONGO_INITDB_ROOT_PASSWORD"),
	}

	client, ctx, cancel, err := mongodb.Connect("mongodb://mongodb/", credential)
	if err != nil {
		panic(err)
	}

	database := client.Database(dbname)

	// TODO: check collections in another file and then use it
	collections := Collections{}
	collections.filters = database.Collection("filters")
	collections.volumes = database.Collection("volumes")
	collections.stocks = database.Collection("stocks")

	// Release resource when the main
	// function is returned.
	defer mongodb.Close(client, ctx, cancel)

	// Ping mongoDB with Ping method
	mongodb.Ping(client, ctx)
}

type Collections struct {
	filters *mongo.Collection
	volumes *mongo.Collection
	stocks  *mongo.Collection
}
