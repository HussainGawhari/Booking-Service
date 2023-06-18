package models

type SeatClass struct {
	ClassName   string  `json:"class_name" gorm:"primaryKey"`
	MinPrice    float64 `json:"min_price"`
	MaxPrice    float64 `json:"max_price"`
	NormalPrice float64 `json:"normal_price"`
}
