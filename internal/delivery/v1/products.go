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
		products.GET("/:id", h.getProductById)

		subcategory := products.Group("/subcategory")
		{
			subcategory.GET("/:id", h.getProductBySubcategoryId)
		}
	}
}

type createProductInp struct {
	Name          string  `json:"name"`
	Price         float64 `json:"price"`
	Quantity      int     `json:"quantity"`
	Description   string  `json:"description"`
	CategoryId    int     `json:"category_id"`
	SubcategoryId int     `json:"subcategory_id"`
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

func (h *Handler) getProductById(c *gin.Context) {
	id, err := processIntParam(c.Param("id"))
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	product, err := h.services.Products.GetById(id)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, product)
}

func (h *Handler) getProductBySubcategoryId(c *gin.Context) {
	id, err := processIntParam(c.Param("id"))
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	product, err := h.services.Products.GetBySubcategoryId(id)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, product)
}

// type updateProductInp struct {
// 	Id            int     `json:"id"`
// 	Name          string  `json:"name"`
// 	Price         float64 `json:"price"`
// 	Quantity      int     `json:"quantity"`
// 	Description   string  `json:"description"`
// 	CategoryId    int     `json:"category_id"`
// 	SubcategoryId int     `json:"subcategory_id"`
// }

// func (h *Handler) updateProduct(c *gin.Context) {
// 	var inp updateProductInp
// 	if err := c.BindJSON(&inp); err != nil {
// 		newResponse(c, http.StatusBadRequest, err.Error())

// 		return
// 	}

// 	product, err := h.services.Products.Update(&domain.Product{
// 		Id:            inp.Id,
// 		Name:          inp.Name,
// 		Price:         inp.Price,
// 		Quantity:      inp.Quantity,
// 		Description:   inp.Description,
// 		CategoryId:    inp.CategoryId,
// 		SubcategoryId: inp.SubcategoryId,
// 	})
// 	if err != nil {
// 		newResponse(c, http.StatusInternalServerError, err.Error())

// 		return
// 	}

// 	c.JSON(http.StatusOK, product)
// }
