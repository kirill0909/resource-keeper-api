package service

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"github.com/kirill0909/resource-keeper-api/models"
	"github.com/kirill0909/resource-keeper-api/pkg/repository"
	"io"
	"os"
)

type UserResourceService struct {
	repo repository.UserResource
}

func NewUserResourceService(repo repository.UserResource) *UserResourceService {
	return &UserResourceService{repo: repo}
}

func (s *UserResourceService) CreateResource(resource models.UserResource) (int, error) {
	if err := encrypt(&resource.ResourcePassword); err != nil {
		return 0, err
	}

	if err := encrypt(&resource.ResourceLogin); err != nil {
		return 0, err
	}

	return s.repo.CreateResource(resource)
}

func (s *UserResourceService) GetAllResources(userId int) ([]models.UserResource, error) {
	return s.repo.GetAllResources(userId)
}

func encrypt(sensitiveData *string) error {
	block, err := aes.NewCipher([]byte(os.Getenv("ENCRYPTION_KEY")))
	if err != nil {
		return err
	}

	cipherText := make([]byte, len(*sensitiveData)+aes.BlockSize)

	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], []byte(*sensitiveData))

	*sensitiveData = hex.EncodeToString(cipherText)
	return nil
}
