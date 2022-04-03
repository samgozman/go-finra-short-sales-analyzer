package stock

import (
	"testing"

	"github.com/samgozman/go-finra-short-sales-analyzer/internal/models/volume"
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

	if got_total != want_total {
		t.Errorf("Expected '%v', but got '%v'", want_total, got_total)
	}
	if got_short != want_short {
		t.Errorf("Expected '%v', but got '%v'", want_short, got_short)
	}
	if got_exempt != want_exempt {
		t.Errorf("Expected '%v', but got '%v'", want_exempt, got_exempt)
	}
}
