package service

import (
	"github.com/kirill0909/resource-keeper-api/models"
	"github.com/kirill0909/resource-keeper-api/pkg/repository"
)

type UserResourceService struct {
	repo repository.UserResource
}

func NewUserResourceService(repo repository.UserResource) *UserResourceService {
	return &UserResourceService{repo: repo}
}

func (s *UserResourceService) CreateResource(resource models.UserResource) (int, error) {
	resource.ResourcePassword = generatePasswordHash(resource.ResourcePassword)
	return s.repo.CreateResource(resource)
}
