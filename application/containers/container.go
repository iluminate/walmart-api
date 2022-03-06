package containers

import (
	"walmart-api/application/handlers"
	"walmart-api/application/services"
)

func ProductHandler() *handlers.ProductHandler {
	return handlers.NewProductHandler(ProductService())
}

func ProductService() services.IProductService {
	return services.NewProductService()
}
