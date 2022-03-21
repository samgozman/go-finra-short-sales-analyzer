package stock

import "fmt"

func CalculateAverages(s []*Stock) {
	// Do some stuff to calculate averages and store them by pointer
	for _, s := range s {
		fmt.Println(s.Ticker)
	}
}

func UpdateMany(s []*Stock) {
	// Update many in DB
}
