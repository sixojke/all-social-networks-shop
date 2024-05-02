package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) initTelegramRoutes(api *gin.RouterGroup) {
	telegram := api.Group("/telegram", h.tgbotIdentity)
	{
		telegram.PUT("/bind", h.bindTelegram)
	}
}

type bindTelegramInp struct {
	AuthCode   string `json:"auth_code" binding:"required"`
	TelegramId int    `json:"telegram_id" binding:"required"`
}

// @Summary Telegram Bind
// @Security TelegramAuth
// @Tags telegram
// @Description bind the telegram account to the account on the site, returns the site user id
// @ModuleID bindTelegram
// @Accept  json
// @Produce  json
// @Param input body bindTelegramInp true "authorization code and telegram for registration in the system"
// @Success 200 {object} idResponse
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /telegram/bind [put]
func (h *Handler) bindTelegram(c *gin.Context) {
	var inp bindTelegramInp
	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	userId, err := h.services.Telegram.Bind(inp.TelegramId, inp.AuthCode)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, idResponse{ID: userId})
}
