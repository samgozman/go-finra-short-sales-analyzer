package stock

import (
	"context"

	"github.com/samgozman/go-finra-short-sales-analyzer/internal/models/volume"
	"github.com/samgozman/go-finra-short-sales-analyzer/pkg/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Calculate average volumes for array of Stock instances
func CalculateAverages(ctx context.Context, db *mongo.Database, lrt int64, stocks *[]Stock) {
	logger.Info("CalculateAverages", "Process started")
	defer logger.Info("CalculateAverages", "Process finished")

	for _, s := range *stocks {
		v := volume.FindLastVolumes(ctx, db, s.ID, 20)
		temp := Stock{ID: s.ID, Ticker: s.Ticker}

		// Check that the volume array is exists and stock was traded during last day
		if len(v) < 1 && v[0].Date.UnixMilli() != lrt {
			// if it's not - make it blanc
			s = temp
			continue
		}

		calcStockVolumeAverages(v, &temp)

		// TODO: Find a way to update all stocks with one pack insert (all at the same time)
		// Update in db
		UpdateOne(ctx, db, temp)
		// Store results in memory for future usage
		s = temp
	}
}

// Update stock by id
func UpdateOne(ctx context.Context, db *mongo.Database, s Stock) {
	update := bson.M{
		"$set": bson.M{
			"shortVolRatioLast":         s.ShortVolRatioLast,
			"shortExemptVolRatioLast":   s.ShortExemptVolRatioLast,
			"shortVolRatio5DAVG":        s.ShortVolRatio5DAVG,
			"shortExemptVolRatio5DAVG":  s.ShortExemptVolRatio5DAVG,
			"shortVolRatio20DAVG":       s.ShortVolRatio20DAVG,
			"shortExemptVolRatio20DAVG": s.ShortExemptVolRatio20DAVG,

			"shortExemptVolLast":   s.ShortExemptVolLast,
			"shortExemptVol5DAVG":  s.ShortExemptVol5DAVG,
			"shortExemptVol20DAVG": s.ShortExemptVol20DAVG,
			"shortVolLast":         s.ShortVolLast,
			"shortVol5DAVG":        s.ShortVol5DAVG,
			"shortVol20DAVG":       s.ShortVol20DAVG,
			"totalVolLast":         s.TotalVolLast,
			"totalVol5DAVG":        s.TotalVol5DAVG,
			"totalVol20DAVG":       s.TotalVol20DAVG,
		},
	}

	_, err := db.Collection("stocks").UpdateOne(
		ctx,
		bson.M{"_id": s.ID},
		update,
	)

	if err != nil {
		logger.Error("UpdateOne", "Error while updating stock "+s.Ticker)
		panic(err)
	}
}

// Calculate groups of averages for a given volume and pass them to the Stock
func calcStockVolumeAverages(v []volume.Volume, s *Stock) {
	// last day (copy just to be able to sort faster without population)
	s.TotalVolLast = v[0].TotalVolume
	s.ShortVolLast = v[0].ShortVolume
	s.ShortExemptVolLast = v[0].ShortExemptVolume
	s.ShortVolRatioLast = float64(v[0].ShortVolume) / float64(v[0].TotalVolume) * 100
	s.ShortExemptVolRatioLast = float64(v[0].ShortExemptVolume) / float64(v[0].TotalVolume) * 100

	// 5 days
	if len(v) >= 5 {
		t5, s5, e5 := avgVolume(v[:5])
		s.TotalVol5DAVG = t5
		s.ShortVol5DAVG = s5
		s.ShortExemptVol5DAVG = e5
		s.ShortVolRatio5DAVG = (s5 / t5) * 100
		s.ShortExemptVolRatio5DAVG = (e5 / t5) * 100
	}

	// 20 days
	if len(v) >= 20 {
		t20, s20, e20 := avgVolume(v)
		s.TotalVol20DAVG = t20
		s.ShortVol20DAVG = s20
		s.ShortExemptVol20DAVG = e20
		s.ShortVolRatio20DAVG = (s20 / t20) * 100
		s.ShortExemptVolRatio20DAVG = (e20 / t20) * 100
	}
}

// Calculate average volumes for a slice
// totalVolume, shortVolume and shortExemptVolume
func avgVolume(vol []volume.Volume) (total float64, short float64, exempt float64) {
	var totalVolume uint64
	var shortVolume uint64
	var shortExemptVolume uint64

	for _, v := range vol {
		totalVolume += v.TotalVolume
		shortVolume += v.ShortVolume
		shortExemptVolume += v.ShortExemptVolume
	}

	arraySize := uint64(len(vol))

	totalVolumeAvg := float64(totalVolume / arraySize)
	shortVolumeAvg := float64(shortVolume / arraySize)
	shortExemptVolumeAvg := float64(shortExemptVolume / arraySize)

	return totalVolumeAvg, shortVolumeAvg, shortExemptVolumeAvg
}
