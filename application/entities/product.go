package entities

type Product struct {
	Id          int    `bson:"id"`
	Brand       string `bson:"brand"`
	Description string `bson:"description"`
	Image       string `bson:"image"`
	Price       int    `bson:"price"`
}
