package models

type Product struct {
	Id          int     `json:"id"`
	Brand       string  `json:"brand"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Price       float64 `json:"price"`
	Discount    float64 `json:"discount,omitempty"`
	Total       float64 `json:"total,omitempty"`
}
