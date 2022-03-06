package storage

import "walmart-api/application/entities"

type IProductStorage interface {
	Find(id string) []entities.Product
}
