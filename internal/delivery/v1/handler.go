package v1

import (
	"fmt"
	"strconv"

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

func processIntParam(param string) (int, error) {
	paramInt, err := strconv.Atoi(param)
	if err != nil {
		return 0, fmt.Errorf("error process int param: %v", err)
	}

	return paramInt, nil
}
