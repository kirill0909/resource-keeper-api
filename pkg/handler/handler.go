package handler

import "github.com/gin-gonic/gin"

type Handler struct{}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signIn)
		auth.POST("/sign-in", h.signUp)
		auth.GET("/ping", h.ping)
	}

	resource := router.Group("/resource")
	{
		resource.POST("/", h.createResource)
		resource.GET("/", h.getAllResource)
		resource.GET("/:id", h.getResourceById)
		resource.PUT("/:id", h.updateResource)
		resource.DELETE("/:id", h.deleteResource)
	}

	return router
}
