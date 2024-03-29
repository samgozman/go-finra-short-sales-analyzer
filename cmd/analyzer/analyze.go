package main

import (
	"fmt"
	"os"

	"github.com/getsentry/sentry-go"
	"github.com/samgozman/go-finra-short-sales-analyzer/internal/models/filter"
	"github.com/samgozman/go-finra-short-sales-analyzer/internal/models/stock"
	"github.com/samgozman/go-finra-short-sales-analyzer/internal/models/volume"
	"github.com/samgozman/go-finra-short-sales-analyzer/pkg/logger"
	"github.com/samgozman/go-finra-short-sales-analyzer/pkg/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Run full analyzer process - from averages update to filters creation
func Run() {
	logger.Info("Run", "The filter update process has been initiated")
	defer logger.Info("Run", "The filter update process has been finished")

	dbname := os.Getenv("MONGODB_NAME")
	dburl := os.Getenv("MONGODB_URL")
	dbport := os.Getenv("MONGODB_PORT")

	credential := options.Credential{
		Username: os.Getenv("MONGO_INITDB_ROOT_USERNAME"),
		Password: os.Getenv("MONGO_INITDB_ROOT_PASSWORD"),
	}

	client, ctx, cancel, err := mongodb.Connect(fmt.Sprintf("mongodb://%s:%s/", dburl, dbport), credential)
	if err != nil {
		sentry.CaptureException(err)
		panic(err)
	}

	database := client.Database(dbname)

	// ! 1. Get all stocks
	cursor, err := database.Collection("stocks").Find(ctx, bson.M{})
	if err != nil {
		logger.Error("Run", "Error while trying to get all stocks from the collection")
		sentry.CaptureException(err)
		panic(err)
	}

	var stArr []stock.Stock
	if err = cursor.All(ctx, &stArr); err != nil {
		logger.Error("Run", "Error while trying to decode each stock item")
		sentry.CaptureException(err)
		panic(err)
	}

	// ! 2. Get latest DB record time (useful to check that the volume is not outdated)
	lrt := volume.LastRecordTime(ctx, database)
	// ! 3. Calculate averages and save them in stocks array by pointer, update in db
	stock.CalculateAverages(ctx, database, lrt, &stArr)
	// ! 4. Drop filters collection
	filter.Drop(ctx, database)
	// ! 5. Pass pointer to a stocks to each filter
	filters := filter.CreateMany(ctx, database, lrt, &stArr)
	// ! 6. Save each filter individually (1 insert transaction for all stocks in 1 filter)
	filter.InsertMany(ctx, database, &filters)

	// Release resource when the main
	// function is returned.
	defer mongodb.Close(ctx, client, cancel)
}
