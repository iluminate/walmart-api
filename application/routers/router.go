package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"walmart-api/application/containers"
	"walmart-api/application/handlers"
)

type httpRouter struct {
	productHandler *handlers.ProductHandler
}

func NewHttpRouter() *httpRouter {
	return &httpRouter{
		productHandler: containers.ProductHandler(),
	}
}

func (httpRouter *httpRouter) Handler() *gin.Engine {
	router := gin.Default()
	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "api is up!")
	})
	router.GET("/api/v1/products/:id", httpRouter.productHandler.Find)
	return router
}
