package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/kirill0909/resource-keeper-api/models"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(email, password string) (models.User, error)
}

type User interface{}

type UserResource interface {
	CreateResource(resource models.UserResource) (int, error)
	GetAllResources(userId int) ([]models.UserResource, error)
	GetById(userId, resourceId int) (models.UserResource, error)
	DeleteResource(userId, resourceId int) error
}

type Repository struct {
	Authorization
	User
	UserResource
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		UserResource:  NewUserResourcePostgres(db),
	}
}
