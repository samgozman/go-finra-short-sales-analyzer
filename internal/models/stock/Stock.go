package stock

import "go.mongodb.org/mongo-driver/bson/primitive"

// DB model that holds ticker symbol and precalculated average values
type Stock struct {
	ID     primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Ticker string             `json:"ticker,omitempty"`

	ShortVolRatioLast         float64 `json:"shortVolRatioLast"`
	ShortExemptVolRatioLast   float64 `json:"shortExemptVolRatioLast"`
	ShortVolRatio5DAVG        float64 `json:"shortVolRatio5DAVG"`
	ShortExemptVolRatio5DAVG  float64 `json:"shortExemptVolRatio5DAVG"`
	ShortVolRatio20DAVG       float64 `json:"shortVolRatio20DAVG"`
	ShortExemptVolRatio20DAVG float64 `json:"shortExemptVolRatio20DAVG"`

	ShortExemptVolLast   float64 `json:"shortExemptVolLast"`
	ShortExemptVol5DAVG  float64 `json:"shortExemptVol5DAVG"`
	ShortExemptVol20DAVG float64 `json:"shortExemptVol20DAVG"`
	ShortVolLast         float64 `json:"shortVolLast"`
	ShortVol5DAVG        float64 `json:"shortVol5DAVG"`
	ShortVol20DAVG       float64 `json:"shortVol20DAVG"`
	TotalVolLast         float64 `json:"totalVolLast"`
	TotalVol5DAVG        float64 `json:"totalVol5DAVG"`
	TotalVol20DAVG       float64 `json:"totalVol20DAVG"`
}
