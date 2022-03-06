package services

import (
	"reflect"
	"testing"
	"walmart-api/application/models"
	"walmart-api/application/services"
	"walmart-api/application/storage"
)

func TestNewProductService(t *testing.T) {
	type args struct {
		productStorage   storage.IProductStorage
		promotionService services.IPromotionService
	}
	tests := []struct {
		name string
		args args
		want *services.ProductService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := services.NewProductService(tt.args.productStorage, tt.args.promotionService); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProductService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewPromotionService(t *testing.T) {
	tests := []struct {
		name string
		want *services.PromotionService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := services.NewPromotionService(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPromotionService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductService_FindBy(t *testing.T) {
	type fields struct {
		productStorage   storage.IProductStorage
		promotionService services.IPromotionService
	}
	type args struct {
		filters map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []models.Product
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &services.ProductService{
				productStorage:   tt.fields.productStorage,
				promotionService: tt.fields.promotionService,
			}
			got, err := service.FindBy(tt.args.filters)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindBy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindBy() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductService_FindById(t *testing.T) {
	type fields struct {
		productStorage   storage.IProductStorage
		promotionService services.IPromotionService
	}
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []models.Product
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &services.ProductService{
				productStorage:   tt.fields.productStorage,
				promotionService: tt.fields.promotionService,
			}
			got, err := service.FindById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindById() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPromotionService_IsPalidrome(t *testing.T) {
	type args struct {
		search string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &services.PromotionService{}
			if got := service.IsPalidrome(tt.args.search); got != tt.want {
				t.Errorf("IsPalidrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPromotionService_ToDiscountPalindrome(t *testing.T) {
	type args struct {
		price float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &services.PromotionService{}
			if got := service.ToDiscountPalindrome(tt.args.price); got != tt.want {
				t.Errorf("ToDiscountPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
