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
func LastRecordTime(ctx context.Context, db *mongo.Database) (m int64) {
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

// Returns 3 arrays of volumes different types and 2 ratios from one volumes array
func SeparateVolumes(volumes *[]Volume) VolumesSeparated {
	var totalVolume []uint64
	var shortVolume []uint64
	var exemptVolume []uint64
	var shortRatio []float32
	var exemptRatio []float32

	for _, v := range *volumes {
		totalVolume = append(totalVolume, v.TotalVolume)
		shortVolume = append(shortVolume, v.ShortVolume)
		exemptVolume = append(exemptVolume, v.ShortExemptVolume)

		shortRatio = append(shortRatio, float32(v.ShortVolume)/float32(v.TotalVolume))
		exemptRatio = append(exemptRatio, float32(v.ShortExemptVolume)/float32(v.TotalVolume))
	}

	res := VolumesSeparated{
		TotalVolume:       totalVolume,
		ShortVolume:       shortVolume,
		ShortExemptVolume: exemptVolume,
		ShortRatio:        shortRatio,
		ExemptRatio:       exemptRatio,
	}

	return res
}
