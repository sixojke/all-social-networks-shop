package v1

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sixojke/internal/domain"
)

func (h *Handler) initProductsRoutes(api *gin.RouterGroup) {
	products := api.Group("/products")
	{
		// products.POST("", h.createProduct)
		// products.GET("/:id", h.getProductById)
		products.GET("", h.productsGetAll)

		// subcategory := products.Group("/subcategory")
		// {
		// 	// subcategory.GET("/:id", h.getProductBySubcategoryId)
		// }
	}
}

// type createProductInp struct {
// 	Name          string  `json:"name"`
// 	Price         float64 `json:"price"`
// 	Quantity      int     `json:"quantity"`
// 	Description   string  `json:"description"`
// 	CategoryId    int     `json:"category_id"`
// 	SubcategoryId int     `json:"subcategory_id"`
// }

// func (h *Handler) createProduct(c *gin.Context) {
// 	var product createProductInp
// 	if err := c.BindJSON(&product); err != nil {
// 		newResponse(c, http.StatusBadRequest, err.Error())

// 		return
// 	}

// 	id, err := h.services.Products.Create(&domain.Product{
// 		Name:          product.Name,
// 		Price:         product.Price,
// 		Quantity:      product.Quantity,
// 		Description:   product.Description,
// 		CategoryId:    product.CategoryId,
// 		SubcategoryId: product.SubcategoryId,
// 	})
// 	if err != nil {
// 		newResponse(c, http.StatusInternalServerError, err.Error())

// 		return
// 	}

// 	c.JSON(http.StatusOK, idResponse{ID: id})
// }

// @Summary User Verify Registration
// @Tags products
// @Description get all products
// @ModuleID allProducts
// @Accept  json
// @Produce  json
// @Param limit query int false "Number of items per page" default(10)
// @Param offset query int false "Offset for pagination" default(0)
// @Success 200 {object} dataResponse
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /products [get]
func (h *Handler) productsGetAll(c *gin.Context) {
	limit, err := processIntParam(c.Query("limit"))
	if err != nil {
		limit = h.config.Pagination.DefaultLimit
	}

	offset, err := processIntParam(c.Query("offset"))
	if err != nil {
		offset = 0
	}

	if limit > h.config.Pagination.MaxLimit {
		limit = h.config.Pagination.MaxLimit
	}

	products, err := h.services.Products.GetAll(limit, offset)
	if err != nil {
		if errors.Is(err, domain.ErrProductsNotFound) {
			newResponse(c, http.StatusOK, err.Error())

			return
		}

		newResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, dataResponse{
		Data:  products,
		Count: int64(len(*products)),
	})
}

// func (h *Handler) getProductById(c *gin.Context) {
// 	id, err := processIntParam(c.Param("id"))
// 	if err != nil {
// 		newResponse(c, http.StatusBadRequest, err.Error())

// 		return
// 	}

// 	product, err := h.services.Products.GetById(id)
// 	if err != nil {
// 		newResponse(c, http.StatusInternalServerError, err.Error())

// 		return
// 	}

// 	c.JSON(http.StatusOK, product)
// }

// func (h *Handler) getProductBySubcategoryId(c *gin.Context) {
// 	id, err := processIntParam(c.Param("id"))
// 	if err != nil {
// 		newResponse(c, http.StatusBadRequest, err.Error())

// 		return
// 	}

// 	product, err := h.services.Products.GetBySubcategoryId(id)
// 	if err != nil {
// 		newResponse(c, http.StatusInternalServerError, err.Error())

// 		return
// 	}

// 	c.JSON(http.StatusOK, product)
// }

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
