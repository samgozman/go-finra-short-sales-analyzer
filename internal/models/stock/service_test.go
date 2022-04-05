package stock

import (
	"testing"

	"github.com/samgozman/go-finra-short-sales-analyzer/internal/models/volume"
	"github.com/samgozman/go-finra-short-sales-analyzer/pkg/tester"
)

func TestAvgVolume(t *testing.T) {
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
	}

	got_total, got_short, got_exempt := avgVolume(volumes)
	var want_total, want_short, want_exempt float64 = 300, 125, 10

	tester.Compare(t, want_total, got_total)
	tester.Compare(t, want_short, got_short)
	tester.Compare(t, want_exempt, got_exempt)
}
