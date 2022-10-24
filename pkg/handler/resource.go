package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kirill0909/resource-keeper-api/models"
	"net/http"
)

func (h *Handler) createResource(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		return
	}

	var input models.UserResource
	input.UID = userId
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	if err = checkEmptyUserResource(input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	resourceId, err := h.service.UserResource.CreateResource(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": resourceId,
	})
}

func (h *Handler) getAllResource(c *gin.Context) {}

func (h *Handler) getResourceById(c *gin.Context) {}

func (h *Handler) updateResource(c *gin.Context) {}

func (h *Handler) deleteResource(c *gin.Context) {}
