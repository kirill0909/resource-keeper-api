package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kirill0909/resource-keeper-api/models"
)

func (h *Handler) signUp(c *gin.Context) {
	var input models.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	if err := checkEmptyValueUser(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

type signInInput struct {
	Email    string `json:"email"    binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := checkEmptyValueSignInInputUser(input.Email, input.Password); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.service.Authorization.GenerateToken(input.Email, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (h *Handler) ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "pong",
	})
}
