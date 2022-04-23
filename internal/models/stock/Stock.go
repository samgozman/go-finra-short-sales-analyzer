package stock

import "go.mongodb.org/mongo-driver/bson/primitive"

// DB model that holds ticker symbol and precalculated average values
type Stock struct {
	ID     primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Ticker string             `bson:"ticker" json:"ticker,omitempty"`

	ShortVolRatioLast         float64 `bson:"shortVolRatioLast" json:"shortVolRatioLast"`
	ShortExemptVolRatioLast   float64 `bson:"shortExemptVolRatioLast" json:"shortExemptVolRatioLast"`
	ShortVolRatio5DAVG        float64 `bson:"shortVolRatio5DAVG" json:"shortVolRatio5DAVG"`
	ShortExemptVolRatio5DAVG  float64 `bson:"shortExemptVolRatio5DAVG" json:"shortExemptVolRatio5DAVG"`
	ShortVolRatio20DAVG       float64 `bson:"shortVolRatio20DAVG" json:"shortVolRatio20DAVG"`
	ShortExemptVolRatio20DAVG float64 `bson:"shortExemptVolRatio20DAVG" json:"shortExemptVolRatio20DAVG"`

	ShortExemptVolLast   uint64  `bson:"shortExemptVolLast" json:"shortExemptVolLast"`
	ShortExemptVol5DAVG  float64 `bson:"shortExemptVol5DAVG" json:"shortExemptVol5DAVG"`
	ShortExemptVol20DAVG float64 `bson:"shortExemptVol20DAVG" json:"shortExemptVol20DAVG"`
	ShortVolLast         uint64  `bson:"shortVolLast" json:"shortVolLast"`
	ShortVol5DAVG        float64 `bson:"shortVol5DAVG" json:"shortVol5DAVG"`
	ShortVol20DAVG       float64 `bson:"shortVol20DAVG" json:"shortVol20DAVG"`
	TotalVolLast         uint64  `bson:"totalVolLast" json:"totalVolLast"`
	TotalVol5DAVG        float64 `bson:"totalVol5DAVG" json:"totalVol5DAVG"`
	TotalVol20DAVG       float64 `bson:"totalVol20DAVG" json:"totalVol20DAVG"`
}
