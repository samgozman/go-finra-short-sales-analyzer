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

			ShortVolGrows5D:     isVolumeGrows(&sv.ShortVolume, 5),
			ShortVolDecreases5D: isVolumeDecreases(&sv.ShortVolume, 5),
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

// ?: This copy-paste is to make code simpler
// TODO: find a better way (without 10x nested if-statements)
// TODO: use generic type (1.18) for volumes

func isVolumeGrows(volumes *[]uint64, daysGrow int) bool {
	if len(*volumes) <= daysGrow {
		return false
	}

	// ! Check volumes with Date - do we need to reverse the order before?
	for i := 1; i < daysGrow+1; i++ {
		isGreaterThanPrev := (*volumes)[i] > (*volumes)[i-1]
		if !isGreaterThanPrev {
			return false
		}
	}

	return true
}

func isVolumeDecreases(volumes *[]uint64, daysGrow int) bool {
	if len(*volumes) <= daysGrow {
		return false
	}

	// ! Check volumes with Date - do we need to reverse the order before?
	for i := 1; i < daysGrow+1; i++ {
		isLesserThanPrev := (*volumes)[i] < (*volumes)[i-1]
		if !isLesserThanPrev {
			return false
		}
	}

	return true
}

// TODO: Replace with generic type. Move this code to pkg
func isRatioGrows(volumes *[]float32, daysGrow int) bool {
	if len(*volumes) <= daysGrow {
		return false
	}

	// ! Check volumes with Date - do we need to reverse the order before?
	for i := 1; i < daysGrow+1; i++ {
		isGreaterThanPrev := (*volumes)[i] > (*volumes)[i-1]
		if !isGreaterThanPrev {
			return false
		}
	}

	return true
}

// TODO: Replace with generic type. Move this code to pkg
func isRatioDecreases(volumes *[]float32, daysGrow int) bool {
	if len(*volumes) <= daysGrow {
		return false
	}

	// ! Check volumes with Date - do we need to reverse the order before?
	for i := 1; i < daysGrow+1; i++ {
		isLesserThanPrev := (*volumes)[i] < (*volumes)[i-1]
		if !isLesserThanPrev {
			return false
		}
	}

	return true
}
