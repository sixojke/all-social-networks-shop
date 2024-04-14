package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sixojke/internal/config"
	v1 "github.com/sixojke/internal/delivery/v1"
	"github.com/sixojke/internal/service"
	"github.com/sixojke/pkg/auth"
)

type Handler struct {
	config       config.Handler
	service      *service.Service
	tokenManager auth.TokenManager
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Init() *gin.Engine {
	router := gin.Default()

	// Init router
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	h.initAPI(router)

	return router
}

func (h *Handler) initAPI(router *gin.Engine) {
	handlerV1 := v1.NewHandler(h.service, h.tokenManager)
	api := router.Group("/api")
	{
		handlerV1.Init(api)
	}
}
