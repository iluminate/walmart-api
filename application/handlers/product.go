package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"walmart-api/application/services"
)

type ProductHandler struct {
	productService services.IProductService
}

func NewProductHandler(productService services.IProductService) *ProductHandler {
	return &ProductHandler{productService: productService}
}

func (handler *ProductHandler) Find(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(http.StatusInternalServerError, "invalid id field")
	}
	products, _ := handler.productService.Find(id)
	c.JSON(http.StatusOK, products)
}
