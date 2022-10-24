package handler

import (
	"errors"
	"strings"

	"github.com/kirill0909/resource-keeper-api/models"
)

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
