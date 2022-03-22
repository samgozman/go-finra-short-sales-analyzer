package volume

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Get number (ms) of the last trading day
func LastDateTime(ctx context.Context, db *mongo.Database) (m int64) {
	// Get latest volume
	findOptions := options.FindOne()
	findOptions.SetSort(bson.D{{"date", -1}})

	var vol Volume
	if err := db.Collection("volumes").FindOne(ctx, bson.M{}, findOptions).Decode(&vol); err != nil {
		log.Fatal(err)
	}

	return vol.Date.UnixMilli()
}

// TODO: make less parameters
// Find last N volume records for the given stockId
func FindLastVolumes(ctx context.Context, db *mongo.Database, stockId primitive.ObjectID, limit int64) []Volume {
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"date", -1}})
	findOptions.SetLimit(limit)

	v, err := db.Collection("volumes").Find(ctx, bson.D{{"_stock_id", stockId}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	var volumes []Volume
	v.All(ctx, &volumes)

	return volumes
}
