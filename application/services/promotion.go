package services

type PromotionService struct{}

func NewPromotionService() *PromotionService {
	return &PromotionService{}
}

func (service *PromotionService) IsPalidrome(search string) bool {
	if search == "" {
		return false
	}
	for i := 0; i < len(search); i++ {
		j := len(search) - 1 - i
		if search[i] != search[j] {
			return false
		}
	}
	return true
}

func (service *PromotionService) ToDiscountPalindrome(price float64) float64 {
	return price - (price * discountRatePalindrome / 100)
}
