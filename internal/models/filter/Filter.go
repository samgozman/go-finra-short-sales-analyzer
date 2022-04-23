package filter

import "go.mongodb.org/mongo-driver/bson/primitive"

// DB model that holds precompiled filters for each stock
type Filter struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	StockId primitive.ObjectID `bson:"_stock_id,omitempty" json:"stock_id"`

	IsNotGarbage bool `json:"isNotGarbage"` // Filter new stocks with no data or incomplete

	ShortVolGrows5D                bool `json:"shortVolGrows5D"`                // Short volume is growing 5 days in a row
	ShortVolDecreases5D            bool `json:"shortVolDecreases5D"`            // Short volume is decreasing 5 days in a row
	ShortVolRatioGrows5D           bool `json:"shortVolRatioGrows5D"`           // Short volume ratio (%) is growing 5 days in a row
	ShortVoRatiolDecreases5D       bool `json:"shortVoRatiolDecreases5D"`       // Short volume ratio (%) is decreasing 5 days in a row
	TotalVolGrows5D                bool `json:"totalVolGrows5D"`                // Total volume is growing 5 days in a row
	TotalVolDecreases5D            bool `json:"totalVolDecreases5D"`            // Total volume is decreasing 5 days in a row
	ShortExemptVolGrows5D          bool `json:"shortExemptVolGrows5D"`          // Short Exempt volume is growing 5 days in a row
	ShortExemptVolDecreases5D      bool `json:"shortExemptVolDecreases5D"`      // Short Exempt volume is decreasing 5 days in a row
	ShortExemptVolRatioGrows5D     bool `json:"shortExemptVolRatioGrows5D"`     // Short Exempt volume ratio is growing 5 days in a row
	ShortExemptVolRatioDecreases5D bool `json:"shortExemptVolRatioDecreases5D"` // Short Exempt volume ratio is decreasing 5 days in a row

	ShortVolGrows3D                bool `json:"shortVolGrows3D"`                // Short volume is growing 3 days in a row
	ShortVolDecreases3D            bool `json:"shortVolDecreases3D"`            // Short volume is decreasing 3 days in a row
	ShortVolRatioGrows3D           bool `json:"shortVolRatioGrows3D"`           // Short volume ratio (%) is growing 3 days in a row
	ShortVoRatiolDecreases3D       bool `json:"shortVoRatiolDecreases3D"`       // Short volume ratio (%) is decreasing 3 days in a row
	TotalVolGrows3D                bool `json:"totalVolGrows3D"`                // Total volume is growing 3 days in a row
	TotalVolDecreases3D            bool `json:"totalVolDecreases3D"`            // Total volume is decreasing 3 days in a row
	ShortExemptVolGrows3D          bool `json:"shortExemptVolGrows3D"`          // Short Exempt volume is growing 3 days in a row
	ShortExemptVolDecreases3D      bool `json:"shortExemptVolDecreases3D"`      // Short Exempt volume is decreasing 3 days in a row
	ShortExemptVolRatioGrows3D     bool `json:"shortExemptVolRatioGrows3D"`     // Short Exempt volume ratio is growing 3 days in a row
	ShortExemptVolRatioDecreases3D bool `json:"shortExemptVolRatioDecreases3D"` // Short Exempt volume ratio is decreasing 3 days in a row

	AbnormalShortlVolGrows          bool `json:"abnormalShortlVolGrows"`          // Abnormal short volume growing
	AbnormalShortVolDecreases       bool `json:"abnormalShortVolDecreases"`       // Abnormal short volume decreasing
	AbnormalTotalVolGrows           bool `json:"abnormalTotalVolGrows"`           // Abnormal total volume growing
	AbnormalTotalVolDecreases       bool `json:"abnormalTotalVolDecreases"`       // Abnormal total volume decreasing
	AbnormalShortExemptVolGrows     bool `json:"abnormalShortExemptVolGrows"`     // Abnormal short exempt volume growing
	AbnormalShortExemptVolDecreases bool `json:"abnormalShortExemptVolDecreases"` // Abnormal short exempt volume decreasing
}
