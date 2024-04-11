package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sixojke/internal/domain"
)

func (h *Handler) initProductsRoutes(api *gin.RouterGroup) {
	products := api.Group("/products")
	{
		products.POST("", h.createProduct)
	}
}

type createProductInp struct {
	Name          string  `json:"name" db:"name"`
	Price         float64 `json:"price" db:"price"`
	Quantity      int     `json:"quantity" db:"quantity"`
	Description   string  `json:"description" db:"description"`
	CategoryId    int     `json:"category_id" db:"category_id"`
	SubcategoryId int     `json:"subcategory_id" db:"subcategory_id"`
}

func (h *Handler) createProduct(c *gin.Context) {
	var product createProductInp
	if err := c.BindJSON(&product); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	id, err := h.services.Products.Create(&domain.Product{
		Name:          product.Name,
		Price:         product.Price,
		Quantity:      product.Quantity,
		Description:   product.Description,
		CategoryId:    product.CategoryId,
		SubcategoryId: product.SubcategoryId,
	})
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, idResponse{ID: id})
}
