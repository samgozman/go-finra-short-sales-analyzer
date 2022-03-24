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
		fi = append(fi, &f)
	}

	_, err := db.Collection("filters").InsertMany(ctx, fi)
	if err != nil {
		fmt.Println("Error while trying to insert new filters")
		panic(err)
	}
}

//analyzer_1  | Iteration 15821
// analyzer_1  | Iteration 15822

// FILTERS
// ? Create filter for each stock

func CreateMany(ctx context.Context, db *mongo.Database, stocks *[]stock.Stock) []Filter {
	var filters []Filter
	lrt := volume.LastRecordTime(ctx, db)

	for _, s := range *stocks {
		f := Filter{
			ID:      primitive.NewObjectID(),
			StockId: s.ID,

			OnTinkoff:    false, // TODO
			IsNotGarbage: isNotGarbageFilter(ctx, db, lrt, s.ID),
		}

		filters = append(filters, f)
	}

	return filters
}

func isNotGarbageFilter(ctx context.Context, db *mongo.Database, lrt int64, stockId primitive.ObjectID) bool {
	volumes := volume.FindLastVolumes(ctx, db, stockId, 5)

	var isConsistent bool = true
	var averageIsAboveMinimum bool = false

	if len(volumes) == 5 && volumes[0].Date.UnixMilli() == lrt {
		var total uint64

		for _, v := range volumes {
			// Check if the volume is not filled up
			if v.TotalVolume == 0 {
				isConsistent = false
			}
			total += v.TotalVolume
		}

		averageIsAboveMinimum = total/5 >= 5000
	}

	if isConsistent && averageIsAboveMinimum {
		return true
	} else {
		return false
	}
}
