package filter

import (
	"context"

	"github.com/samgozman/go-finra-short-sales-analyzer/internal/models/stock"
	"github.com/samgozman/go-finra-short-sales-analyzer/internal/models/volume"
	"github.com/samgozman/go-finra-short-sales-analyzer/pkg/logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Drop filters collection
func Drop(ctx context.Context, db *mongo.Database) {
	err := db.Collection("filters").Drop(ctx)

	if err != nil {
		logger.Error("Drop", "Error while trying to drop Filters collection")
		panic(err)
	}
}

// Insert array of Filters entities into Collection
func InsertMany(ctx context.Context, db *mongo.Database, filters *[]Filter) {
	logger.Info("InsertMany", "Updating filters process started")
	defer logger.Info("InsertMany", "Updating filters process finished")

	// TODO: Find a way to create interface[] of filters from the start
	// Convert struct into interface
	var fi []interface{}
	for _, f := range *filters {
		fi = append(fi, f)
	}

	_, err := db.Collection("filters").InsertMany(ctx, fi)
	if err != nil {
		logger.Error("InsertMany", "Error while trying to insert new filters")
		panic(err)
	}
}

// FILTERS
// ? Create filter for each stock

// Create many filters objects
func CreateMany(ctx context.Context, db *mongo.Database, lrt int64, stocks *[]stock.Stock) []Filter {
	logger.Info("CreateMany", "Process started")
	defer logger.Info("CreateMany", "Process finished")

	var filters []Filter

	for _, s := range *stocks {
		// TODO: Ja-ja, giant N+1
		volumes := volume.FindLastVolumes(ctx, db, s.ID, 5)
		currentLatestRecord := volumes[0].Date.UnixMilli()

		// Reverse volumes for filters usage
		volumes = volume.Reverse(&volumes)
		sv := volume.SeparateVolumes(&volumes)

		f := Filter{
			ID:      primitive.NewObjectID(),
			StockId: s.ID,

			IsNotGarbage: isNotGarbageFilter(lrt, currentLatestRecord, &sv.TotalVolume),

			// TODO: Refactor idea - return number of grow days in a row and compare it with 5
			ShortVolGrows5D:                isGrowing(&sv.ShortVolume, 5),
			ShortVolDecreases5D:            isDeclining(&sv.ShortVolume, 5),
			ShortVolRatioGrows5D:           isGrowing(&sv.ShortRatio, 5),
			ShortVoRatiolDecreases5D:       isDeclining(&sv.ShortRatio, 5),
			TotalVolGrows5D:                isGrowing(&sv.TotalVolume, 5),
			TotalVolDecreases5D:            isDeclining(&sv.TotalVolume, 5),
			ShortExemptVolGrows5D:          isGrowing(&sv.ShortExemptVolume, 5),
			ShortExemptVolDecreases5D:      isDeclining(&sv.ShortExemptVolume, 5),
			ShortExemptVolRatioGrows5D:     isGrowing(&sv.ExemptRatio, 5),
			ShortExemptVolRatioDecreases5D: isDeclining(&sv.ExemptRatio, 5),

			ShortVolGrows3D:                isGrowing(&sv.ShortVolume, 3),
			ShortVolDecreases3D:            isDeclining(&sv.ShortVolume, 3),
			ShortVolRatioGrows3D:           isGrowing(&sv.ShortRatio, 3),
			ShortVoRatiolDecreases3D:       isDeclining(&sv.ShortRatio, 3),
			TotalVolGrows3D:                isGrowing(&sv.TotalVolume, 3),
			TotalVolDecreases3D:            isDeclining(&sv.TotalVolume, 3),
			ShortExemptVolGrows3D:          isGrowing(&sv.ShortExemptVolume, 3),
			ShortExemptVolDecreases3D:      isDeclining(&sv.ShortExemptVolume, 3),
			ShortExemptVolRatioGrows3D:     isGrowing(&sv.ExemptRatio, 3),
			ShortExemptVolRatioDecreases3D: isDeclining(&sv.ExemptRatio, 3),

			AbnormalShortlVolGrows:          isAbnormalGrowth(s.ShortVol20DAVG, s.ShortVolLast),
			AbnormalShortVolDecreases:       isAbnormalDecline(s.ShortVol20DAVG, s.ShortVolLast),
			AbnormalTotalVolGrows:           isAbnormalGrowth(s.TotalVol20DAVG, s.TotalVolLast),
			AbnormalTotalVolDecreases:       isAbnormalDecline(s.TotalVol20DAVG, s.TotalVolLast),
			AbnormalShortExemptVolGrows:     isAbnormalGrowth(s.ShortExemptVol20DAVG, s.ShortExemptVolLast),
			AbnormalShortExemptVolDecreases: isAbnormalDecline(s.ShortExemptVol20DAVG, s.ShortExemptVolLast),
		}

		filters = append(filters, f)
	}

	return filters
}

func isNotGarbageFilter(lastRecordTime int64, currentRecordTime int64, totalVolumes *[]uint64) bool {
	var isConsistent bool = true
	var averageIsAboveMinimum bool = false

	if len(*totalVolumes) >= 5 && currentRecordTime == lastRecordTime {
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

func isGrowing[T numeric](volumes *[]T, daysGrow int) bool {
	if len(*volumes) < daysGrow {
		return false
	}

	for i := 1; i < daysGrow; i++ {
		isGreaterThanPrev := (*volumes)[i] > (*volumes)[i-1]
		if !isGreaterThanPrev {
			return false
		}
	}

	return true
}

func isDeclining[T numeric](volumes *[]T, daysGrow int) bool {
	if len(*volumes) < daysGrow {
		return false
	}

	for i := 1; i < daysGrow; i++ {
		isLesserThanPrev := (*volumes)[i] < (*volumes)[i-1]
		if !isLesserThanPrev {
			return false
		}
	}

	return true
}

// Abnormal volume => more than triple the 20d average
func isAbnormalGrowth(average float64, current uint64) bool {
	multiplier := float64(current) / average

	if multiplier >= 3 {
		return true
	} else {
		return false
	}
}

// Abnormal volume => more than triple the 20d average
func isAbnormalDecline(average float64, current uint64) bool {
	multiplier := average / float64(current)

	if multiplier >= 3 {
		return true
	} else {
		return false
	}
}

type numeric interface {
	int | int64 | uint | float32 | float64 | uint64
}
