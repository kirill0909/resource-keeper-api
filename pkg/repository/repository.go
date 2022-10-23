package repository

type Authorization interface{}

type User interface{}

type UserResource interface{}

type Repository struct {
	Authorization
	User
	UserResource
}

func NewRepository() *Repository {
	return &Repository{}
}
