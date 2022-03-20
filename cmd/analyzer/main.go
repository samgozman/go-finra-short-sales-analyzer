package main

import (
	"fmt"
	"os"

	"github.com/samgozman/go-finra-short-sales-analyzer/internal/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	dbname := os.Getenv("MONGODB_NAME")

	credential := options.Credential{
		Username: os.Getenv("MONGO_INITDB_ROOT_USERNAME"),
		Password: os.Getenv("MONGO_INITDB_ROOT_PASSWORD"),
	}

	client, ctx, cancel, err := mongodb.Connect("mongodb://mongodb/"+dbname, credential)
	if err != nil {
		panic(err)
	}

	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		fmt.Println(err)

	}
	fmt.Println(databases)

	// Release resource when the main
	// function is returned.
	defer mongodb.Close(client, ctx, cancel)

	// Ping mongoDB with Ping method
	mongodb.Ping(client, ctx)
}
