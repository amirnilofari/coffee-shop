package models

type Drink struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}