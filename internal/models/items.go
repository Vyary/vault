package models

type Price struct {
	Value  float64 `json:"value"`
	Type   string  `json:"type"`
	Listed int     `json:"listed"`
}

type UniquesDTO struct {
	ItemID int    `json:"id"`
	Name   string `json:"name"`
	Base   string `json:"base"`
	Image  string `json:"image"`
	Price  Price  `json:"price"`
}
