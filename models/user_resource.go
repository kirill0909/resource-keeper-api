package models

type UserResource struct {
	Id               int    `json:"-"`
	UID              int    `json:"-"`
	ResourceName     string `json:"resource_name"`
	ResourceLogin    string `json:"resource_login"`
	ResourcePassword string `json:"resource_password"`
	DateCreation     string `json:"-"`
	LastUpdate       string `json:"-"`
}
