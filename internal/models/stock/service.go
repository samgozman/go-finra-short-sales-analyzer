package stock

import (
	"context"
	"fmt"

	"github.com/samgozman/go-finra-short-sales-analyzer/internal/models/volume"
	"go.mongodb.org/mongo-driver/mongo"
)

func CalculateAverages(ctx context.Context, db *mongo.Database, s *[]Stock) {
	// Do some stuff to calculate averages and store them by pointer
	lastRecordTime := volume.LastDateTime(ctx, db)
	fmt.Println("Last date:", lastRecordTime)

	// Get last date from volume service
	for _, s := range *s {
		v := volume.FindLastVolumes(ctx, db, s.ID, 20)
		fmt.Println("Ticker:", s.Ticker, " last volumes:", v)

		// Check that the volume array is exists and stock was traded during last day
		if len(v) > 1 && v[0].Date.UnixMilli() == lastRecordTime {
			// ! Calc averages for last day, 5, 20
		}
	}
}

func UpdateMany(s []*Stock) {
	// Update many in DB
}
