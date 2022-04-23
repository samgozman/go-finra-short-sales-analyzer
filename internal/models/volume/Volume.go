package volume

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// DB model that holds basic short volume data from FINRA
type Volume struct {
	ID                primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	StockId           primitive.ObjectID `bson:"_stock_id" json:"stock_id,omitempty"`
	Date              time.Time          `bson:"date" json:"date"`
	ShortVolume       uint64             `bson:"shortVolume" json:"shortVolume"`
	ShortExemptVolume uint64             `bson:"shortExemptVolume" json:"shortExemptVolume"`
	TotalVolume       uint64             `bson:"totalVolume" json:"totalVolume"`
}

// List of volumes data usabale in filter conditions
type VolumesSeparated struct {
	ShortVolume       []uint64
	ShortExemptVolume []uint64
	TotalVolume       []uint64
	ShortRatio        []float32 // Array of short volumes divided by total volume
	ExemptRatio       []float32 // Array of short exempt volumes divided by total volume
}
