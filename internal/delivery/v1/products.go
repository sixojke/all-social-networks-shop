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

	categories := api.Group("/categories")
	{
		categories.GET("", h.categoriesAll)
		categories.GET("/:id", h.subcategoriesById)
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

// @Summary Products with filters
// @Tags products
// @Description Get all products with pagination and filters
// @ModuleID allProducts
// @Accept  json
// @Produce  json
// @Param limit query int false "Number of items per page" default(10) maximum(100)
// @Param offset query int false "Offset for pagination" default(0)
// @Param category_id query int false "category_id"
// @Param subcategory_id query int false "subcategory_id"
// @Param is_available query int false "Product availability: enter 1 if true, 0 if false" default(0)
// @Param sort_price query string false "sort price: enter asc or desc"
// @Param sort_defect query string false "sort defect: enter asc or desc"
// @Success 200 {object} paginationResponse
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

	categoryId, _ := processIntParam(c.Query("category_id"))

	subcategoryId, _ := processIntParam(c.Query("subcategory_id"))

	isAvailable, _ := processIntParam(c.Query("is_available"))

	sortPrice := c.Query("sort_price")

	sortDefect := c.Query("sort_defect")

	if limit > h.config.Pagination.MaxLimit {
		limit = h.config.Pagination.MaxLimit
	}

	products, err := h.services.Products.GetAll(&domain.ProductFilters{
		Limit:         limit,
		Offset:        offset,
		CategoryId:    categoryId,
		SubcategoryId: subcategoryId,
		IsAvailable:   isAvailable,
		SortPrice:     sortPrice,
		SortDefect:    sortDefect,
	})
	if err != nil {
		if errors.Is(err, domain.ErrProductsNotFound) {
			newResponse(c, http.StatusOK, err.Error())

			return
		}

		newResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, paginationResponse{
		Pagination: products,
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

// @Summary Get all categories
// @Tags categories
// @Description get all categories
// @ModuleID allCategories
// @Accept  json
// @Produce  json
// @Success 200 {object} dataResponse
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /categories [get]
func (h *Handler) categoriesAll(c *gin.Context) {
	categories, err := h.services.Category.GetCategories()
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, dataResponse{
		Data:  categories,
		Count: int64(len(*categories)),
	})
}

// @Summary Get subcategories
// @Tags subcategories
// @Description  get subcategories by category_id
// @ModuleID subcategoriesById
// @Accept  json
// @Produce  json
// @Param category_id path int true "category id"
// @Success 200 {object} dataResponse
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /categories/{category_id} [get]
func (h *Handler) subcategoriesById(c *gin.Context) {
	categoryId, err := processIntParam(c.Param("id"))
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	subcategories, err := h.services.Category.GetSubcategories(categoryId)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, dataResponse{
		Data:  subcategories,
		Count: int64(len(*subcategories)),
	})
}
