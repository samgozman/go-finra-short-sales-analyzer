package filter

import (
	"testing"
	"time"
)

func TestIsNotGarbageFilter(t *testing.T) {
	t.Run("Should return true if data is consistent, avg above min and record time is equal", func(t *testing.T) {
		// Params
		lrt := time.Now().UnixMilli()
		crt := lrt
		totalVolumes := []uint64{5600, 6600, 7500, 5000, 9000}

		want := true
		got := isNotGarbageFilter(lrt, crt, &totalVolumes)

		if want != got {
			t.Errorf("Expected '%v', but got '%v' with params: lrt '%v', crt '%v', total '%v'", want, got, lrt, crt, totalVolumes)
		}
	})
	t.Run("Should return false if data is inconsistent", func(t *testing.T) {
		// Params
		lrt := time.Now().UnixMilli()
		crt := lrt
		totalVolumes := []uint64{10600, 10600, 0, 15000, 29000}

		want := false
		got := isNotGarbageFilter(lrt, crt, &totalVolumes)

		if want != got {
			t.Errorf("Expected '%v', but got '%v' with params: lrt '%v', crt '%v', total '%v'", want, got, lrt, crt, totalVolumes)
		}
	})
	t.Run("Should return false if data is not enough", func(t *testing.T) {
		// Params
		lrt := time.Now().UnixMilli()
		crt := lrt
		totalVolumes := []uint64{10600, 16600, 17500, 15000}

		want := false
		got := isNotGarbageFilter(lrt, crt, &totalVolumes)

		if want != got {
			t.Errorf("Expected '%v', but got '%v' with params: lrt '%v', crt '%v', total '%v'", want, got, lrt, crt, totalVolumes)
		}
	})
	t.Run("Should return false if volume is below minimum", func(t *testing.T) {
		// Params
		lrt := time.Now().UnixMilli()
		crt := lrt
		totalVolumes := []uint64{600, 700, 2000, 350, 200}

		want := false
		got := isNotGarbageFilter(lrt, crt, &totalVolumes)

		if want != got {
			t.Errorf("Expected '%v', but got '%v' with params: lrt '%v', crt '%v', total '%v'", want, got, lrt, crt, totalVolumes)
		}
	})
	t.Run("Should return false if last record time and current are not equal", func(t *testing.T) {
		// This checks if passed volumes for the current period refer
		// to the latest DB entry (that the volumes are not outdated)

		// Params
		lrt := time.Now().UnixMilli()
		crt := time.Now().UnixMilli() + 100
		totalVolumes := []uint64{600, 700, 2000, 350, 200}

		want := false
		got := isNotGarbageFilter(lrt, crt, &totalVolumes)

		if want != got {
			t.Errorf("Expected '%v', but got '%v' with params: lrt '%v', crt '%v', total '%v'", want, got, lrt, crt, totalVolumes)
		}
	})
}

func TestIsGrowing(t *testing.T) {
	t.Run("Should return true if volumes are growing one by one", func(t *testing.T) {
		totalVolumes := []uint64{3000, 3001, 3002, 3003, 3005}

		want := true
		got := isGrowing(&totalVolumes, 5)

		if want != got {
			t.Errorf("Expected '%v', but got '%v' with total '%v'", want, got, totalVolumes)
		}
	})
	t.Run("Should return false if volumes are not growing", func(t *testing.T) {
		totalVolumes := []uint64{3000, 2000, 1000, 500, 800}

		want := false
		got := isGrowing(&totalVolumes, 5)

		if want != got {
			t.Errorf("Expected '%v', but got '%v' with total '%v'", want, got, totalVolumes)
		}
	})
	t.Run("Should return false if volumes len are less than 'daysGrow'", func(t *testing.T) {
		totalVolumes := []uint64{3000, 3001, 3002, 3003}

		want := false
		got := isGrowing(&totalVolumes, 5)

		if want != got {
			t.Errorf("Expected '%v', but got '%v' with total '%v'", want, got, totalVolumes)
		}
	})
}

func TestIsDeclining(t *testing.T) {
	t.Run("Should return true if volumes are declining one by one", func(t *testing.T) {
		totalVolumes := []uint64{3000, 2999, 2998, 2997, 2996}

		want := true
		got := isDeclining(&totalVolumes, 5)

		if want != got {
			t.Errorf("Expected '%v', but got '%v' with total '%v'", want, got, totalVolumes)
		}
	})
	t.Run("Should return false if volumes are not declining", func(t *testing.T) {
		totalVolumes := []uint64{3000, 3001, 3002, 3003, 3005}

		want := false
		got := isDeclining(&totalVolumes, 5)

		if want != got {
			t.Errorf("Expected '%v', but got '%v' with total '%v'", want, got, totalVolumes)
		}
	})
	t.Run("Should return false if volumes len are less than 'daysGrow'", func(t *testing.T) {
		totalVolumes := []uint64{3000, 2999, 2998, 2997}

		want := false
		got := isDeclining(&totalVolumes, 5)

		if want != got {
			t.Errorf("Expected '%v', but got '%v' with total '%v'", want, got, totalVolumes)
		}
	})
}

func TestIsAbnormalGrowth(t *testing.T) {
	t.Run("Should return true if current volume is above avg 3 times", func(t *testing.T) {
		var average float64 = 30.611
		var current uint64 = 95

		want := true
		got := isAbnormalGrowth(average, current)

		if want != got {
			t.Errorf("Expected '%v', but got '%v' with average '%v' and current '%v", want, got, average, current)
		}
	})
	t.Run("Should return false if current volume is not above avg 3 times", func(t *testing.T) {
		var average float64 = 30.611
		var current uint64 = 21

		want := false
		got := isAbnormalGrowth(average, current)

		if want != got {
			t.Errorf("Expected '%v', but got '%v' with average '%v' and current '%v", want, got, average, current)
		}
	})
}

func TestIsAbnormaDecline(t *testing.T) {
	t.Run("Should return true if current volume is below avg 3 times", func(t *testing.T) {
		var average float64 = 30.611
		var current uint64 = 5

		want := true
		got := isAbnormaDecline(average, current)

		if want != got {
			t.Errorf("Expected '%v', but got '%v' with average '%v' and current '%v", want, got, average, current)
		}
	})
	t.Run("Should return false if current volume is not below avg 3 times", func(t *testing.T) {
		var average float64 = 30.611
		var current uint64 = 21

		want := false
		got := isAbnormaDecline(average, current)

		if want != got {
			t.Errorf("Expected '%v', but got '%v' with average '%v' and current '%v", want, got, average, current)
		}
	})
}
