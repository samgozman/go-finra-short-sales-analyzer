package main

import (
	"log"
	"os"

	"github.com/samgozman/go-finra-short-sales-analyzer/internal/models/stock"
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

	client, ctx, cancel, err := mongodb.Connect("mongodb://mongodb/", credential)
	if err != nil {
		panic(err)
	}

	database := client.Database(dbname)

	// ! 1. Get all stocks
	cursor, err := database.Collection("stocks").Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	var stArr []stock.Stock
	if err = cursor.All(ctx, &stArr); err != nil {
		log.Fatal(err)
	}

	// ! 2. Calculate averages and save them in stocks array by pointer
	stock.CalculateAverages(ctx, database, &stArr)
	// ! 3. Save all stocks with calculated fields with one query (update many)
	// stock.UpdateMany(stArr)
	// ! 4. Drop filters collection
	// ! 5. Pass pointer to a stocks to each filter
	// ! 6. Save each filter individually (1 insert transaction for all stocks in 1 filter)

	// Release resource when the main
	// function is returned.
	defer mongodb.Close(client, ctx, cancel)

	// Ping mongoDB with Ping method
	mongodb.Ping(client, ctx)
}
