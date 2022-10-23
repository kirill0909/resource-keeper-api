package service

import (
	"github.com/kirill0909/resource-keeper-api/models"
	"github.com/kirill0909/resource-keeper-api/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
}

type User interface{}

type UserResource interface{}

type Service struct {
	Authorization
	User
	UserResource
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
	}
}
