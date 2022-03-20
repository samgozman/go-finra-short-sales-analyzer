package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// DB model that holds basic short volume data from FINRA
type Volume struct {
	ID                primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	StockId           primitive.ObjectID `bson:"_stock_id" json:"stock_id,omitempty"`
	Date              time.Time          `json:"date"`
	ShortVolume       uint64             `json:"shortVolume"`
	ShortExemptVolume uint64             `json:"shortExemptVolume"`
	TotalVolume       uint64             `json:"totalVolume"`
}
