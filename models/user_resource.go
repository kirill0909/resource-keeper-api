package models

import (
	"errors"
	"strings"
)

type UserResource struct {
	Id               int    `json:"id"                db:"id"`
	UID              int    `json:"user_id"           db:"user_id"`
	ResourceName     string `json:"resource_name"     db:"resource_name"          binding:"required"`
	ResourceLogin    string `json:"resource_login"    db:"resource_login_enc"     binding:"required"`
	ResourcePassword string `json:"resource_password" db:"resource_password_enc"  binding:"required"`
	DateCreation     string `json:"date_creation"     db:"date_creation"`
	LastUpdate       string `json:"last_update"       db:"last_update"`
}

type UserResourceUpdate struct {
	ResourceName     *string `json:"resource_name"`
	ResourceLogin    *string `json:"resource_login"`
	ResourcePassword *string `json:"resource_password"`
}

func (u *UserResourceUpdate) Validate() error {
	if (u.ResourceName != nil && len(strings.TrimSpace(*u.ResourceName)) == 0) ||
		(u.ResourceLogin != nil && len(strings.TrimSpace(*u.ResourceLogin)) == 0) ||
		(u.ResourcePassword != nil && len(strings.TrimSpace(*u.ResourcePassword)) == 0) {
		return errors.New("update structure has no values")
	}

	return nil
}
