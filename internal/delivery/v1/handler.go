package v1

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sixojke/internal/config"
	"github.com/sixojke/internal/service"
	"github.com/sixojke/pkg/auth"
)

type Handler struct {
	config       config.Handler
	services     *service.Service
	tokenManager auth.TokenManager
}

func NewHandler(config config.Handler, services *service.Service, tokenManager auth.TokenManager) *Handler {
	return &Handler{
		config:       config,
		services:     services,
		tokenManager: tokenManager,
	}
}

func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("v1")
	{
		h.initTelegramRoutes(v1)
		h.initAdminRoutes(v1)
		h.initUsersRoutes(v1)
		h.initProductsRoutes(v1)
		h.initReferralSystemRoutes(v1)
	}
}

func (h *Handler) getLimitAndOffset(c *gin.Context) (limit int, offset int) {
	limit, err := processIntParam(c.Query("limit"))
	if err != nil {
		limit = h.config.Pagination.DefaultLimit
	}

	page, err := processIntParam(c.Query("page"))
	if err != nil {
		page = 1
	}

	if limit > h.config.Pagination.MaxLimit {
		limit = h.config.Pagination.MaxLimit
	}

	return limit, page*limit - limit
}

func processIntParam(param string) (int, error) {
	paramInt, err := strconv.Atoi(param)
	if err != nil {
		return 0, fmt.Errorf("error process int param: %v", err)
	}

	return paramInt, nil
}
