package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"walmart-api/application/services"
)

type ProductHandler struct {
	productService services.IProductService
}

func NewProductHandler(productService services.IProductService) *ProductHandler {
	return &ProductHandler{productService: productService}
}

func (handler *ProductHandler) Find(c *gin.Context) {
	id := c.Param("id")
	products := handler.productService.Find(id)
	c.JSON(http.StatusOK, products)
}
