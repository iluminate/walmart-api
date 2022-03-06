package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProductHandler struct {
}

func NewProductHandler() *ProductHandler {
	return &ProductHandler{}
}

func (handler *ProductHandler) Find(c *gin.Context) {
	id := c.Param("id")
	c.String(http.StatusOK, "id: %s", id)
}
