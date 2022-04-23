package filter

import "go.mongodb.org/mongo-driver/bson/primitive"

// DB model that holds precompiled filters for each stock
type Filter struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	StockId primitive.ObjectID `bson:"_stock_id,omitempty" json:"stock_id"`

	IsNotGarbage bool `bson:"isNotGarbage" json:"isNotGarbage"` // Filter new stocks with no data or incomplete

	ShortVolGrows5D                bool `bson:"shortVolGrows5D" json:"shortVolGrows5D"`                               // Short volume is growing 5 days in a row
	ShortVolDecreases5D            bool `bson:"shortVolDecreases5D" json:"shortVolDecreases5D"`                       // Short volume is decreasing 5 days in a row
	ShortVolRatioGrows5D           bool `bson:"shortVolRatioGrows5D" json:"shortVolRatioGrows5D"`                     // Short volume ratio (%) is growing 5 days in a row
	ShortVoRatiolDecreases5D       bool `bson:"shortVoRatiolDecreases5D" json:"shortVoRatiolDecreases5D"`             // Short volume ratio (%) is decreasing 5 days in a row
	TotalVolGrows5D                bool `bson:"totalVolGrows5D" json:"totalVolGrows5D"`                               // Total volume is growing 5 days in a row
	TotalVolDecreases5D            bool `bson:"totalVolDecreases5D" json:"totalVolDecreases5D"`                       // Total volume is decreasing 5 days in a row
	ShortExemptVolGrows5D          bool `bson:"shortExemptVolGrows5D" json:"shortExemptVolGrows5D"`                   // Short Exempt volume is growing 5 days in a row
	ShortExemptVolDecreases5D      bool `bson:"shortExemptVolDecreases5D" json:"shortExemptVolDecreases5D"`           // Short Exempt volume is decreasing 5 days in a row
	ShortExemptVolRatioGrows5D     bool `bson:"shortExemptVolRatioGrows5D" json:"shortExemptVolRatioGrows5D"`         // Short Exempt volume ratio is growing 5 days in a row
	ShortExemptVolRatioDecreases5D bool `bson:"shortExemptVolRatioDecreases5D" json:"shortExemptVolRatioDecreases5D"` // Short Exempt volume ratio is decreasing 5 days in a row

	ShortVolGrows3D                bool `bson:"shortVolGrows3D" json:"shortVolGrows3D"`                               // Short volume is growing 3 days in a row
	ShortVolDecreases3D            bool `bson:"shortVolDecreases3D" json:"shortVolDecreases3D"`                       // Short volume is decreasing 3 days in a row
	ShortVolRatioGrows3D           bool `bson:"shortVolRatioGrows3D" json:"shortVolRatioGrows3D"`                     // Short volume ratio (%) is growing 3 days in a row
	ShortVoRatiolDecreases3D       bool `bson:"shortVoRatiolDecreases3D" json:"shortVoRatiolDecreases3D"`             // Short volume ratio (%) is decreasing 3 days in a row
	TotalVolGrows3D                bool `bson:"totalVolGrows3D" json:"totalVolGrows3D"`                               // Total volume is growing 3 days in a row
	TotalVolDecreases3D            bool `bson:"totalVolDecreases3D" json:"totalVolDecreases3D"`                       // Total volume is decreasing 3 days in a row
	ShortExemptVolGrows3D          bool `bson:"shortExemptVolGrows3D" json:"shortExemptVolGrows3D"`                   // Short Exempt volume is growing 3 days in a row
	ShortExemptVolDecreases3D      bool `bson:"shortExemptVolDecreases3D" json:"shortExemptVolDecreases3D"`           // Short Exempt volume is decreasing 3 days in a row
	ShortExemptVolRatioGrows3D     bool `bson:"shortExemptVolRatioGrows3D" json:"shortExemptVolRatioGrows3D"`         // Short Exempt volume ratio is growing 3 days in a row
	ShortExemptVolRatioDecreases3D bool `bson:"shortExemptVolRatioDecreases3D" json:"shortExemptVolRatioDecreases3D"` // Short Exempt volume ratio is decreasing 3 days in a row

	AbnormalShortlVolGrows          bool `bson:"abnormalShortlVolGrows" json:"abnormalShortlVolGrows"`                   // Abnormal short volume growing
	AbnormalShortVolDecreases       bool `bson:"abnormalShortVolDecreases" json:"abnormalShortVolDecreases"`             // Abnormal short volume decreasing
	AbnormalTotalVolGrows           bool `bson:"abnormalTotalVolGrows" json:"abnormalTotalVolGrows"`                     // Abnormal total volume growing
	AbnormalTotalVolDecreases       bool `bson:"abnormalTotalVolDecreases" json:"abnormalTotalVolDecreases"`             // Abnormal total volume decreasing
	AbnormalShortExemptVolGrows     bool `bson:"abnormalShortExemptVolGrows" json:"abnormalShortExemptVolGrows"`         // Abnormal short exempt volume growing
	AbnormalShortExemptVolDecreases bool `bson:"abnormalShortExemptVolDecreases" json:"abnormalShortExemptVolDecreases"` // Abnormal short exempt volume decreasing
}
