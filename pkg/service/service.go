package service

import (
	"github.com/kirill0909/resource-keeper-api/models"
	"github.com/kirill0909/resource-keeper-api/pkg/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(token string) (int, error)
}

type User interface{}

type UserResource interface {
	CreateResource(resource models.UserResource) (int, error)
	GetAllResources(userId int) ([]models.UserResource, error)
	GetById(userId, resourceId int) (models.UserResource, error)
	UpdateResource(userId, resourceId int, input models.UserResourceUpdate) error
	DeleteResource(userId, resourceId int) (int, error)
}

type Service struct {
	Authorization
	User
	UserResource
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		UserResource:  NewUserResourceService(repo.UserResource),
	}
}
