package services

import "walmart-api/application/models"

const discountRatePalindrome = 50

type IProductService interface {
	FindById(id int64) ([]models.Product, error)
	FindBy(filters map[string]string) ([]models.Product, error)
}

type IPromotionService interface {
	IsPalidrome(search string) bool
	ToDiscountPalindrome(price float64) float64
}
