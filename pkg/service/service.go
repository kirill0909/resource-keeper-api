package service

import "github.com/kirill0909/resource-keeper-api/pkg/repository"

type Authorization interface{}

type User interface{}

type UserResource interface{}

type Service struct {
	Authorization
	User
	UserResource
}

func NewService(repo *repository.Repository) *Service {
	return &Service{}
}
