package volume

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// DB model that holds basic short volume data from FINRA
type Volume struct {
	ID                primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	StockId           primitive.ObjectID `bson:"_stock_id" json:"stock_id,omitempty"`
	Date              time.Time          `json:"date"`
	ShortVolume       float64            `json:"shortVolume"`
	ShortExemptVolume float64            `json:"shortExemptVolume"`
	TotalVolume       float64            `json:"totalVolume"`
}
