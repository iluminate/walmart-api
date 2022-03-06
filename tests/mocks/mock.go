package mocks

import (
	"github.com/stretchr/testify/mock"
	"walmart-api/application/entities"
)

type MockProductStorage struct {
	mock.Mock
}

func (m *MockProductStorage) FindById(id int64) ([]entities.Product, error) {
	args := m.Called(id)
	return args.Get(0).([]entities.Product), args.Error(1)
}

func (m *MockProductStorage) FindBy(filters map[string]string) ([]entities.Product, error) {
	args := m.Called(filters)
	return args.Get(0).([]entities.Product), args.Error(1)
}

type MockPromotionService struct {
	mock.Mock
}

func (m *MockPromotionService) IsPalidrome(search string) bool {
	args := m.Called(search)
	return args.Bool(0)
}

func (m *MockPromotionService) ToDiscountPalindrome(price float64) float64 {
	args := m.Called(price)
	return args.Get(0).(float64)
}
