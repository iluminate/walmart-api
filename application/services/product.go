package services

import (
	"walmart-api/application/models"
)

type ProductService struct{}

func NewProductService() *ProductService {
	return &ProductService{}
}

func (service *ProductService) Find(id string) []models.Product {
	return []models.Product{}
}
