package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/sixojke/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("v1")
	{
		h.initProductsRoutes(v1)
	}
}
