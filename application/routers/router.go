package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type httpRouter struct {
}

func NewHttpRouter() *httpRouter {
	return &httpRouter{}
}

func (httpRouter *httpRouter) Handler() *gin.Engine {
	router := gin.Default()
	router.GET("/api/v1/products/:id", func(context *gin.Context) {
		context.String(http.StatusOK, "")
	})
	return router
}
