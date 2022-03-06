package storage

import "walmart-api/application/entities"

const collectionProducts = "products"

type IProductStorage interface {
	FindById(id int64) ([]entities.Product, error)
	FindBy(filters map[string]string) ([]entities.Product, error)
}
