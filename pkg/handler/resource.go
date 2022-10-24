package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createResource(c *gin.Context) {
	id, _ := c.Get("userCtx")
	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})

}

func (h *Handler) getAllResource(c *gin.Context) {}

func (h *Handler) getResourceById(c *gin.Context) {}

func (h *Handler) updateResource(c *gin.Context) {}

func (h *Handler) deleteResource(c *gin.Context) {}
