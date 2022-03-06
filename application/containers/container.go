package containers

import "walmart-api/application/handlers"

func ProductHandler() *handlers.ProductHandler {
	return handlers.NewProductHandler()
}
