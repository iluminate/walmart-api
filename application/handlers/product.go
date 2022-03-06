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

func (handler *ProductHandler) FindById(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(http.StatusInternalServerError, "invalid field id")
		return
	}
	products, err := handler.productService.FindById(id)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, products)
}

func (handler *ProductHandler) FindBy(c *gin.Context) {
	filters := make(map[string]string, 0)
	for k, v := range c.Request.URL.Query() {
		if len(v) == 1 && len(v[0]) != 0 {
			filters[k] = v[0]
			break
		}
	}
	products, err := handler.productService.FindBy(filters)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, products)
}
