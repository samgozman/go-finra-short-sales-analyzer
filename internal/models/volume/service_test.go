package volume

import (
	"reflect"
	"testing"
)

func TestSeparateVolumes(t *testing.T) {
	volumes := []Volume{
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

	got := SeparateVolumes(&volumes)
	want := VolumesSeparated{
		ShortVolume:       []uint64{100, 150},
		ShortExemptVolume: []uint64{5, 15},
		TotalVolume:       []uint64{200, 400},
		ShortRatio:        []float32{0.5, 0.375},
		ExemptRatio:       []float32{0.025, 0.0375},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected '%v', but got '%v'", want, got)
	}
}

func TestReverse(t *testing.T) {
	volumes := []Volume{
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

	got := Reverse(&volumes)
	want := []Volume{
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
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected '%v', but got '%v'", want, got)
	}
}
