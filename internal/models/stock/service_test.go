package stock

import (
	"testing"

	"github.com/samgozman/go-finra-short-sales-analyzer/internal/models/volume"
	"github.com/samgozman/go-finra-short-sales-analyzer/pkg/tester"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestAvgVolume(t *testing.T) {
	volumes := []volume.Volume{
		{
			ShortVolume:       102,
			ShortExemptVolume: 5,
			TotalVolume:       200,
		},
		{
			ShortVolume:       150,
			ShortExemptVolume: 15,
			TotalVolume:       401,
		},
	}

	got_total, got_short, got_exempt := avgVolume(volumes)
	var want_total, want_short, want_exempt float64 = 300.5, 126, 10

	tester.Compare(t, want_total, got_total)
	tester.Compare(t, want_short, got_short)
	tester.Compare(t, want_exempt, got_exempt)
}

func TestCalcStockVolumeAverages(t *testing.T) {
	volumes := []volume.Volume{
		{
			ShortVolume:       100,
			ShortExemptVolume: 5,
			TotalVolume:       200,
		},
		{
			ShortVolume:       150,
			ShortExemptVolume: 15,
			TotalVolume:       400,
		},
		{
			ShortVolume:       100,
			ShortExemptVolume: 5,
			TotalVolume:       200,
		},
		{
			ShortVolume:       150,
			ShortExemptVolume: 15,
			TotalVolume:       400,
		}, {
			ShortVolume:       100,
			ShortExemptVolume: 5,
			TotalVolume:       200,
		},
		{
			ShortVolume:       150,
			ShortExemptVolume: 15,
			TotalVolume:       400,
		}, {
			ShortVolume:       100,
			ShortExemptVolume: 5,
			TotalVolume:       200,
		},
		{
			ShortVolume:       150,
			ShortExemptVolume: 15,
			TotalVolume:       400,
		}, {
			ShortVolume:       100,
			ShortExemptVolume: 5,
			TotalVolume:       200,
		},
		{
			ShortVolume:       150,
			ShortExemptVolume: 15,
			TotalVolume:       400,
		}, {
			ShortVolume:       100,
			ShortExemptVolume: 5,
			TotalVolume:       200,
		},
		{
			ShortVolume:       150,
			ShortExemptVolume: 15,
			TotalVolume:       400,
		}, {
			ShortVolume:       100,
			ShortExemptVolume: 5,
			TotalVolume:       200,
		},
		{
			ShortVolume:       150,
			ShortExemptVolume: 15,
			TotalVolume:       400,
		}, {
			ShortVolume:       100,
			ShortExemptVolume: 5,
			TotalVolume:       200,
		},
		{
			ShortVolume:       150,
			ShortExemptVolume: 15,
			TotalVolume:       400,
		}, {
			ShortVolume:       100,
			ShortExemptVolume: 5,
			TotalVolume:       200,
		},
		{
			ShortVolume:       150,
			ShortExemptVolume: 15,
			TotalVolume:       400,
		}, {
			ShortVolume:       100,
			ShortExemptVolume: 5,
			TotalVolume:       200,
		},
		{
			ShortVolume:       150,
			ShortExemptVolume: 15,
			TotalVolume:       400,
		},
	}

	id := primitive.NewObjectID()
	ticker := "TEST"

	got := Stock{ID: id, Ticker: ticker}
	calcStockVolumeAverages(volumes, &got)

	want := Stock{
		ID:     id,
		Ticker: ticker,

		ShortVolRatioLast:         50,
		ShortExemptVolRatioLast:   2.5,
		ShortVolRatio5DAVG:        42.857142857142854,
		ShortExemptVolRatio5DAVG:  3.214285714285714,
		ShortVolRatio20DAVG:       41.66666666666667,
		ShortExemptVolRatio20DAVG: 3.3333333333333335,

		ShortExemptVolLast:   5,
		ShortExemptVol5DAVG:  9,
		ShortExemptVol20DAVG: 10,
		ShortVolLast:         100,
		ShortVol5DAVG:        120,
		ShortVol20DAVG:       125,
		TotalVolLast:         200,
		TotalVol5DAVG:        280,
		TotalVol20DAVG:       300,
	}

	tester.StructCompare(t, want, got)
}
