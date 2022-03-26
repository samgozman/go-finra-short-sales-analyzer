package filter

import (
	"context"
	"fmt"

	"github.com/samgozman/go-finra-short-sales-analyzer/internal/models/stock"
	"github.com/samgozman/go-finra-short-sales-analyzer/internal/models/volume"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func Drop(ctx context.Context, db *mongo.Database) {
	err := db.Collection("filters").Drop(ctx)

	if err != nil {
		fmt.Println("Error while trying to drop Filters collection")
		panic(err)
	}
}

// Insert array of Filters entities into Collection
func InsertMany(ctx context.Context, db *mongo.Database, filters *[]Filter) {
	// TODO: Find a way to create interface[] of filters from the start
	// Convert struct into interface
	var fi []interface{}
	for _, f := range *filters {
		fi = append(fi, f)
	}

	_, err := db.Collection("filters").InsertMany(ctx, fi)
	if err != nil {
		fmt.Println("Error while trying to insert new filters")
		panic(err)
	}

	fmt.Println("Filters successfully updated!")
}

// FILTERS
// ? Create filter for each stock

func CreateMany(ctx context.Context, db *mongo.Database, stocks *[]stock.Stock) []Filter {
	var filters []Filter
	lrt := volume.LastRecordTime(ctx, db)

	for _, s := range *stocks {
		// TODO: Ja-ja, giant N+1
		volumes := volume.FindLastVolumes(ctx, db, s.ID, 20)
		currentLatestRecord := volumes[0].Date.UnixMilli()
		sv := volume.SeparateVolumes(&volumes)

		f := Filter{
			ID:      primitive.NewObjectID(),
			StockId: s.ID,

			OnTinkoff:    false, // TODO
			IsNotGarbage: isNotGarbageFilter(lrt, currentLatestRecord, &sv.TotalVolume),
		}

		filters = append(filters, f)
	}

	return filters
}

func isNotGarbageFilter(lastRecordTime int64, curentRecordTime int64, totalVolumes *[]uint64) bool {
	var isConsistent bool = true
	var averageIsAboveMinimum bool = false

	if len(*totalVolumes) >= 5 && curentRecordTime == lastRecordTime {
		var total uint64

		for i := 0; i < 5; i++ {
			// Check if the volume is not filled up
			if (*totalVolumes)[i] == 0 {
				isConsistent = false
			}
			total += (*totalVolumes)[i]
		}

		averageIsAboveMinimum = total/5 >= 5000
	}

	if isConsistent && averageIsAboveMinimum {
		return true
	} else {
		return false
	}
}
