package services

import "walmart-api/application/models"

type IProductService interface {
	Find(id int64) ([]models.Product, error)
}
