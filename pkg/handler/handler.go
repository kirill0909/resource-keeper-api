package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kirill0909/resource-keeper-api/pkg/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		auth.GET("/ping", h.ping)
	}

	resource := router.Group("/resource", h.userIdentity)
	{
		resource.POST("/", h.createResource)
		resource.GET("/", h.getAllResources)
		resource.GET("/:id", h.getResourceById)
		resource.PUT("/:id", h.updateResource)
		resource.DELETE("/:id", h.deleteResource)
	}

	return router
}
