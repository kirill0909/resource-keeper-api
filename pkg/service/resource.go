package service

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"github.com/kirill0909/resource-keeper-api/models"
	"github.com/kirill0909/resource-keeper-api/pkg/repository"
	"io"
	"os"
	"strings"
)

type UserResourceService struct {
	repo repository.UserResource
}

func NewUserResourceService(repo repository.UserResource) *UserResourceService {
	return &UserResourceService{repo: repo}
}

func (s *UserResourceService) CreateResource(resource models.UserResource) (int, error) {
	if err := encrypt([]byte(os.Getenv("ENCRYPTION_KEY")), &resource.ResourcePassword); err != nil {
		return 0, err
	}

	if err := encrypt([]byte(os.Getenv("ENCRYPTION_KEY")), &resource.ResourceLogin); err != nil {
		return 0, err
	}

	return s.repo.CreateResource(resource)
}

func (s *UserResourceService) GetAllResources(userId int) ([]models.UserResource, error) {
	resources, err := s.repo.GetAllResources(userId)
	if err != nil {
		return nil, err
	}
	for index := range resources {
		if err := decrypt([]byte(os.Getenv("ENCRYPTION_KEY")), &resources[index].ResourceLogin); err != nil {
			return nil, err
		}
		if err := decrypt([]byte(os.Getenv("ENCRYPTION_KEY")), &resources[index].ResourcePassword); err != nil {
			return nil, err
		}
	}

	return resources, nil
}

func (s *UserResourceService) GetById(userId, resourceId int) (models.UserResource, error) {
	resource, err := s.repo.GetById(userId, resourceId)
	if err != nil {
		return models.UserResource{}, err
	}

	if err := decrypt([]byte(os.Getenv("ENCRYPTION_KEY")), &resource.ResourceLogin); err != nil {
		return models.UserResource{}, err
	}

	if err := decrypt([]byte(os.Getenv("ENCRYPTION_KEY")), &resource.ResourcePassword); err != nil {
		return models.UserResource{}, err
	}

	return resource, nil
}

func (s *UserResourceService) DeleteResource(userId, resourceId int) (int, error) {
	return s.repo.DeleteResource(userId, resourceId)
}

func (s *UserResourceService) UpdateResource(userId, resourceId int, input models.UserResourceUpdate) error {
	if err := input.Validate(); err != nil {
		return err
	}

	if input.ResourceLogin != nil && len(strings.TrimSpace(*input.ResourceLogin)) != 0 {
		if err := encrypt([]byte(os.Getenv("ENCRYPTION_KEY")), input.ResourceLogin); err != nil {
			return err
		}
	}

	if input.ResourcePassword != nil && len(strings.TrimSpace(*input.ResourcePassword)) != 0 {
		if err := encrypt([]byte(os.Getenv("ENCRYPTION_KEY")), input.ResourcePassword); err != nil {
			return err
		}
	}

	return s.repo.UpdateResource(userId, resourceId, input)
}

func encrypt(key []byte, sensitiveData *string) error {
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	cipherText := make([]byte, len([]byte(*sensitiveData))+aes.BlockSize)

	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], []byte(*sensitiveData))

	*sensitiveData = base64.RawStdEncoding.EncodeToString(cipherText)
	return nil
}

func decrypt(key []byte, secure *string) error {
	cipherText, err := base64.RawStdEncoding.DecodeString(*secure)

	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	if len(cipherText) < aes.BlockSize {
		return errors.New("cipherText block size is too short!")
	}

	iv := (cipherText)[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, []byte(iv))
	stream.XORKeyStream(cipherText, cipherText)

	*secure = string(cipherText)
	return nil
}
