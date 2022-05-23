package models

type Drink struct {
	ID    uint    `json:"id" gorm:"primary_key"`
	Name  string  `json:"name"`
	Type  string  `json:"type"`
	Price float64 `json:"price"`
}
