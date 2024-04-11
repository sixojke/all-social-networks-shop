package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/sixojke/internal/config"
	"github.com/sixojke/internal/service"
)

type Handler struct {
	config  config.Handler
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Init() *gin.Engine {
	router := gin.Default()

	return router
}
