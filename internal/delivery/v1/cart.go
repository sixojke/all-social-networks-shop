package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sixojke/internal/domain"
)

func (h *Handler) initCartRoutes(user *gin.RouterGroup) {
	cart := user.Group("/cart")
	{
		cart.GET("", h.cartByUserId)
		cart.PUT("/product", h.cartSetQuantity)
	}
}

// @Summary User Cart
// @Security UsersAuth
// @Tags user
// @Description get products from cart
// @ModuleID cartByUserId
// @Accept  json
// @Produce  json
// @Success 200 {object} dataResponse
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /user/cart [get]
func (h *Handler) cartByUserId(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	cart, err := h.services.Cart.GetById(userId)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, dataResponse{Data: cart})
}

type cartSetQuantityInp struct {
	ProductId int `json:"product_id" binding:"required"`
	Quantity  int `json:"quantity" binding:"required"`
}

// @Summary  Update product quantity
// @Security UsersAuth
// @Tags user
// @Description update product quantity
// @ModuleID cartSetQuantity
// @Accept  json
// @Produce  json
// @Param input body cartSetQuantityInp true "product quantity"
// @Success 200 {object} response
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /user/cart/product [put]
func (h *Handler) cartSetQuantity(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	var inp cartSetQuantityInp
	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	if err := h.services.Cart.SetQuantity(&domain.CartSetQuantityInp{
		UserId:    userId,
		ProductId: inp.ProductId,
		Quantity:  inp.Quantity,
	}); err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, response{Message: "success"})
}
