package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kirill0909/resource-keeper-api/models"
	"net/http"
	"strconv"
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

type getAllResourcesResponse struct {
	Data []models.UserResource `json:"data"`
}

func (h *Handler) getAllResources(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		return
	}

	resources, err := h.service.UserResource.GetAllResources(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllResourcesResponse{
		Data: resources,
	})

}

func (h *Handler) getResourceById(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		return
	}

	resourceId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	resource, err := h.service.UserResource.GetById(userId, resourceId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"resource": resource})
}

func (h *Handler) updateResource(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		return
	}

	resourceId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input models.UserResourceUpdate
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.UserResource.UpdateResource(userId, resourceId, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{Status: "ok"})
}

func (h *Handler) deleteResource(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		return
	}

	resourceId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.service.UserResource.DeleteResource(userId, resourceId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{Status: "ok"})
}
