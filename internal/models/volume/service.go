package volume

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
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
