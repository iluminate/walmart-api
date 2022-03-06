package services

import (
	"walmart-api/application/models"
	"walmart-api/application/storage"
)

type ProductService struct {
	productStorage storage.IProductStorage
}

func NewProductService(productStorage storage.IProductStorage) *ProductService {
	return &ProductService{productStorage: productStorage}
}

func (service *ProductService) Find(id string) []models.Product {
	products := service.productStorage.Find(id)
	var result []models.Product
	for _, product := range products {
		result = append(result, models.Product{
			Id:          product.Id,
			Brand:       product.Brand,
			Description: product.Description,
			Image:       product.Image,
			Price:       product.Price,
		})
	}
	return result
}
