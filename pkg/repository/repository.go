package repository

import "github.com/jmoiron/sqlx"

type Authorization interface{}

type User interface{}

type UserResource interface{}

type Repository struct {
	Authorization
	User
	UserResource
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
