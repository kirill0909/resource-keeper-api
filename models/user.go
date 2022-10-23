package models

type User struct {
	Id           int    `json:"-"        db:"id"`
	Name         string `jsojn:"name"    binding:"required"`
	Email        string `json:"email"    binding:"required"`
	Password     string `json:"password" binding:"required"`
	DateCreation string `json:"-"`
	LastUpdate   string `json:"-"`
}
