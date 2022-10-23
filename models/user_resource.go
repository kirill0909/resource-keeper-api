package models

type UserResource struct {
	Id               int    `json:"-"`
	UID              int    `json:"-"`
	ResourceName     string `json:"resource_name"     binding:"required"`
	ResourceLogin    string `json:"resource_login"    binding:"required"`
	ResourcePassword string `json:"resource_password" binding:"required"`
	DateCreation     string `json:"-"`
	LastUpdate       string `json:"-"`
}
