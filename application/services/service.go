package services

import "walmart-api/application/models"

type IProductService interface {
	Find(id string) []models.Product
}
