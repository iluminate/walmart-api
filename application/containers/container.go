package containers

import (
	"walmart-api/application/handlers"
	"walmart-api/application/services"
	"walmart-api/application/storage"
)

func ProductHandler() *handlers.ProductHandler {
	return handlers.NewProductHandler(ProductService())
}

func ProductService() services.IProductService {
	return services.NewProductService(ProductStorage())
}

func ProductStorage() storage.IProductStorage {
	return storage.NewProductStorage()
}
