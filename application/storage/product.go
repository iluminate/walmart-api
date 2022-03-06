package storage

import (
	"walmart-api/application/entities"
)

type ProductStorage struct{}

func NewProductStorage() *ProductStorage {
	return &ProductStorage{}
}

func (storage *ProductStorage) Find(id string) []entities.Product {
	var products []entities.Product
	return products
}
