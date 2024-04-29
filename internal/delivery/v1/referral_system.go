package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) initReferralSystemRoutes(api *gin.RouterGroup) {
	referralSystem := api.Group("/referral-system")
	{
		referralSystem.POST("/visitor", h.referralSystemAddVisitor)
	}
}

// @Summary Referral System Add Visitor
// @Tags referral-system
// @Description add a visitor using a referral code
// @ModuleID referralSystemAddVisitor
// @Accept  json
// @Produce  json
// @Param referral_code query string false "referral code"
// @Success 200 {object} response
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /referral-system/visitor [post]
func (h *Handler) referralSystemAddVisitor(c *gin.Context) {
	referralCode := c.Query("referral_code")

	if referralCode == "" {
		c.JSON(http.StatusOK, response{Message: "success"})

		return
	}

	if err := h.services.ReferralSystem.AddVisitor(referralCode); err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, response{Message: "success"})
}
