package services

import (
	"walmart-api/application/models"
	"walmart-api/application/storage"
)

type ProductService struct {
	productStorage   storage.IProductStorage
	promotionService IPromotionService
}

func NewProductService(productStorage storage.IProductStorage, promotionService IPromotionService) *ProductService {
	return &ProductService{productStorage: productStorage, promotionService: promotionService}
}

func (service *ProductService) FindById(id int64) ([]models.Product, error) {
	products, err := service.productStorage.FindById(id)
	if err != nil {
		return nil, err
	}
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
	return result, nil
}

func (service *ProductService) FindBy(filters map[string]string) ([]models.Product, error) {
	isPromotionPalindrome := false
	for _, v := range filters {
		if service.promotionService.IsPalidrome(v) {
			isPromotionPalindrome = true
		}
	}
	products, err := service.productStorage.FindBy(filters)
	if err != nil {
		return nil, err
	}
	var result []models.Product
	for _, product := range products {
		var discount, total float64
		if isPromotionPalindrome {
			discount = discountRatePalindrome
			total = service.promotionService.ToDiscountPalindrome(product.Price)
		}
		result = append(result, models.Product{
			Id:          product.Id,
			Brand:       product.Brand,
			Description: product.Description,
			Image:       product.Image,
			Price:       product.Price,
			Discount:    discount,
			Total:       total,
		})
	}
	return result, nil
}
