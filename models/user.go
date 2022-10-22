package models

type User struct {
	Id           int    `json:"-"`
	Name         string `jsojn:"name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	DateCreation string `json:"-"`
	LastUpdate   string `json:"-"`
}
