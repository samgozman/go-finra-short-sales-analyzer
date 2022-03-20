package models

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

	ShortExemptVolLast   uint64 `json:"shortExemptVolLast"`
	ShortExemptVol5DAVG  uint64 `json:"shortExemptVol5DAVG"`
	ShortExemptVol20DAVG uint64 `json:"shortExemptVol20DAVG"`
	ShortVolLast         uint64 `json:"shortVolLast"`
	ShortVol5DAVG        uint64 `json:"shortVol5DAVG"`
	ShortVol20DAVG       uint64 `json:"shortVol20DAVG"`
	TotalVolLast         uint64 `json:"totalVolLast"`
	TotalVol5DAVG        uint64 `json:"totalVol5DAVG"`
	TotalVol20DAVG       uint64 `json:"totalVol20DAVG"`
}
