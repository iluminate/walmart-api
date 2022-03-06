package storage

import "walmart-api/application/entities"

const collectionProducts = "products"

type IProductStorage interface {
	Find(id int64) ([]entities.Product, error)
}
