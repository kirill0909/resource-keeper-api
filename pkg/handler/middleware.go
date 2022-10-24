package handler

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kirill0909/resource-keeper-api/models"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	log.Println(header)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	if len(strings.TrimSpace(headerParts[1])) == 0 {
		newErrorResponse(c, http.StatusUnauthorized, "token is empty")
		return
	}

	userId, err := h.service.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set("userCtx", userId)
}

func checkEmptyValueUser(user *models.User) error {
	if len(strings.TrimSpace(user.Name)) == 0 ||
		len(strings.TrimSpace(user.Email)) == 0 ||
		len(strings.TrimSpace(user.Password)) == 0 {
		return errors.New("invalid input body")
	}

	return nil
}

func checkEmptyValueSignInInputUser(email, password string) error {
	if len(strings.TrimSpace(email)) == 0 ||
		len(strings.TrimSpace(password)) == 0 {
		return errors.New("invalid input body")
	}

	return nil
}
