package models

type Product struct {
	Id          int    `json:"id"`
	Brand       string `json:"brand"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Price       int    `json:"price"`
}
