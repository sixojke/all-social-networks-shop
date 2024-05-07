package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) initCartRoutes(user *gin.RouterGroup) {
	cart := user.Group("/cart")
	{
		cart.GET("", h.cartByUserId)
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
