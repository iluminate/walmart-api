package containers

import (
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
	"walmart-api/application/handlers"
	"walmart-api/application/services"
	"walmart-api/application/storage"
	"walmart-api/helpers/database"
)

var syncMongoHelper sync.Once
var mongoHelper *database.MongoHelper

func ProductHandler() *handlers.ProductHandler {
	return handlers.NewProductHandler(ProductService())
}

func ProductService() services.IProductService {
	return services.NewProductService(ProductStorage())
}

func ProductStorage() storage.IProductStorage {
	return storage.NewProductStorage(MongoHelper())
}

func MongoHelper() *database.MongoHelper {
	syncMongoHelper.Do(func() {
		mongoHelper = database.NewMongoHelper(mongoConfig())
		go mongoHelper.OpenConnection()
	})
	return mongoHelper
}

func mongoConfig() *database.MongoConfig {
	clientOptions := options.Client().ApplyURI("mongodb://localhost/admin")
	clientOptions.Auth = &options.Credential{
		Username: "productListUser",
		Password: "productListPassword",
	}
	clientOptions.SetAppName("walmart-api")
	clientOptions.SetMaxPoolSize(100)
	clientOptions.SetMinPoolSize(50)
	return &database.MongoConfig{
		ClientOptions: clientOptions,
		Database:      "promotions",
	}
}
