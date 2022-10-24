package models

type UserResource struct {
	Id               int    `json:"id"                db:"id"`
	UID              int    `json:"user_id"           db:"user_id"`
	ResourceName     string `json:"resource_name"     db:"resource_name"          binding:"required"`
	ResourceLogin    string `json:"resource_login"    db:"resource_login"         binding:"required"`
	ResourcePassword string `json:"resource_password" db:"resource_password_hash" binding:"required"`
	DateCreation     string `json:"date_creation"     db:"date_creation"`
	LastUpdate       string `json:"last_update"       db:"last_update"`
}
