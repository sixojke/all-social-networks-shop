package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/sixojke/docs"
	"github.com/sixojke/internal/config"
	v1 "github.com/sixojke/internal/delivery/v1"
	"github.com/sixojke/internal/service"
	"github.com/sixojke/pkg/auth"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	config       config.Handler
	service      *service.Service
	tokenManager auth.TokenManager
}

func NewHandler(config config.Handler, service *service.Service, tokenManager auth.TokenManager) *Handler {
	return &Handler{
		config:       config,
		service:      service,
		tokenManager: tokenManager,
	}
}

func (h *Handler) Init() *gin.Engine {
	router := gin.Default()

	router.Use(
		corsMiddleware,
	)

	router.GET("/api/v1/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Init router
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	h.initAPI(router)

	return router
}

func (h *Handler) initAPI(router *gin.Engine) {
	handlerV1 := v1.NewHandler(h.config, h.service, h.tokenManager)
	api := router.Group("/api")
	{
		handlerV1.Init(api)
	}
}
