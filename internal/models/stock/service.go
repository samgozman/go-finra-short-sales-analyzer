package stock

import (
	"context"

	"github.com/samgozman/go-finra-short-sales-analyzer/internal/models/volume"
	"go.mongodb.org/mongo-driver/mongo"
)

// Calculate average volumes for array of Stock instances
func CalculateAverages(ctx context.Context, db *mongo.Database, stocks *[]Stock) {
	// Get last date from volume service
	lastRecordTime := volume.LastDateTime(ctx, db)

	for _, s := range *stocks {
		v := volume.FindLastVolumes(ctx, db, s.ID, 20)
		temp := Stock{ID: s.ID, Ticker: s.Ticker}

		// Check that the volume array is exists and stock was traded during last day
		if len(v) < 1 && v[0].Date.UnixMilli() != lastRecordTime {
			// if it's not - make it blanc
			s = temp
			continue
		}

		// last day (copy just to be able to sort faster without population)
		temp.TotalVolLast = v[0].TotalVolume
		temp.ShortVolLast = v[0].ShortVolume
		temp.ShortExemptVolLast = v[0].ShortExemptVolume
		temp.ShortVolRatioLast = float64(v[0].ShortVolume/v[0].TotalVolume) * 100
		temp.ShortExemptVolRatioLast = float64(v[0].ShortExemptVolume/v[0].TotalVolume) * 100

		// 5 days
		if len(v) >= 5 {
			t5, s5, e5 := avgVolume(v[:5])
			temp.TotalVol5DAVG = t5
			temp.ShortVol5DAVG = s5
			temp.ShortExemptVol5DAVG = e5
			temp.ShortVolRatio5DAVG = (s5 / t5) * 100
			temp.ShortExemptVolRatio5DAVG = (e5 / t5) * 100
		}

		// 20 days
		if len(v) >= 20 {
			t20, s20, e20 := avgVolume(v)
			temp.TotalVol20DAVG = t20
			temp.ShortVol20DAVG = s20
			temp.ShortExemptVol20DAVG = e20
			temp.ShortVolRatio20DAVG = (s20 / t20) * 100
			temp.ShortExemptVolRatio20DAVG = (e20 / t20) * 100
		}

		// Store
		s = temp
	}
}

func UpdateMany(s []*Stock) {
	// Update many in DB
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