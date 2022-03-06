package services

import (
	"reflect"
	"testing"
	"walmart-api/application/entities"
	"walmart-api/application/models"
	"walmart-api/application/services"
	"walmart-api/application/storage"
	"walmart-api/tests/mocks"
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
		{name: "create_instance_of_product_service", args: args{
			productStorage:   nil,
			promotionService: nil,
		}, want: &services.ProductService{}},
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
		{name: "create_instance_of_promotion_service", want: &services.PromotionService{}},
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

	mockProductStorage := new(mocks.MockProductStorage)
	mockPromotionService := new(mocks.MockPromotionService)

	mockProductStorage.On("FindBy", map[string]string{"brand": "ada"}).Return([]entities.Product{}, nil)
	mockProductStorage.On("FindBy", map[string]string{"brand": "apa"}).Return([]entities.Product{
		{Id: 1, Brand: "papa", Description: "", Image: "", Price: float64(230)}}, nil)
	mockPromotionService.On("IsPalidrome", "ada").Return(true)
	mockPromotionService.On("IsPalidrome", "apa").Return(true)
	mockPromotionService.On("ToDiscountPalindrome", float64(230)).Return(float64(115))

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
		{name: "search_without_filters_no_match", fields: fields{
			productStorage:   mockProductStorage,
			promotionService: mockPromotionService,
		}, args: args{filters: map[string]string{"brand": "ada"}}, want: nil, wantErr: false},
		{name: "search_without_filters_with_results", fields: fields{
			productStorage:   mockProductStorage,
			promotionService: mockPromotionService,
		}, args: args{filters: map[string]string{"brand": "apa"}}, want: []models.Product{
			{
				Id:          1,
				Brand:       "papa",
				Description: "",
				Image:       "",
				Price:       float64(230),
				Discount:    float64(50),
				Total:       float64(115),
			}}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &services.ProductService{
				ProductStorage:   tt.fields.productStorage,
				PromotionService: tt.fields.promotionService,
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

	mockProductStorage := new(mocks.MockProductStorage)

	mockProductStorage.On("FindById", int64(1)).Return([]entities.Product{}, nil)
	mockProductStorage.On("FindById", int64(2)).Return([]entities.Product{
		{Id: 1, Brand: "", Description: "", Image: "", Price: float64(20)},
	}, nil)

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
		{name: "search_by_id_exist", fields: fields{
			productStorage:   mockProductStorage,
			promotionService: nil,
		}, args: args{id: 2}, want: []models.Product{
			{Id: 1, Brand: "", Description: "", Image: "", Price: float64(20)},
		}, wantErr: false},
		{name: "search_by_id_no_exist", fields: fields{
			productStorage:   mockProductStorage,
			promotionService: nil,
		}, args: args{id: 1}, want: nil, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &services.ProductService{
				ProductStorage:   tt.fields.productStorage,
				PromotionService: tt.fields.promotionService,
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
		{name: "verify_palindrome_true", args: args{"abba"}, want: true},
		{name: "verify_palindrome_false", args: args{"ssds"}, want: false},
		{name: "verify_palindrome_with_word_empty", args: args{""}, want: false},
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
		{name: "apply_discount_palindrome_success", args: args{price: float64(100)}, want: float64(50)},
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
